package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day 8 - Start")
	data := loadData("day8part1.txt")
	fmt.Println("Part 1 -", part1(data, 6, 50))
	fmt.Println("Part 2 -")
	part2(data, 6, 50)
}

// Command with arguments A and B
type Command struct {
	name CommandName
	A, B int
}

// CommandName enum
type CommandName int

// Command type
const (
	RotateRow    CommandName = iota
	RotateColumn             = iota
	Rect                     = iota
)

func nextInt(str string, i int) (int, int) {
	for ; i < len(str) && !unicode.IsDigit(rune(str[i])); i++ {
	}
	x := 0
	for ; i < len(str) && unicode.IsDigit(rune(str[i])); i++ {
		x = x*10 + int(str[i]) - '0'
	}
	return x, i
}

func loadData(filename string) []Command {
	var data []Command

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		cmd := parseCommand(scanner.Text())
		data = append(data, cmd)

	}
	return data
}

func parseCommand(str string) Command {
	var cmd Command
	switch {
	case strings.HasPrefix(str, "rect"):
		cmd.name = Rect
	case strings.HasPrefix(str, "rotate row"):
		cmd.name = RotateRow
	case strings.HasPrefix(str, "rotate column"):
		cmd.name = RotateColumn
	}
	i := 0
	cmd.A, i = nextInt(str, i)
	cmd.B, i = nextInt(str, i+1)
	return cmd
}

func part1(data []Command, heigth, width int) int {
	screen := make([][]bool, heigth)
	for i := range screen {
		screen[i] = make([]bool, width)
	}

	for _, cmd := range data {
		runCommand(&screen, cmd)
	}

	return countLitPixels(screen)
}

func runCommand(screen *[][]bool, cmd Command) {
	switch cmd.name {
	case Rect:
		runRect(screen, cmd.A, cmd.B)
	case RotateRow:
		runRotateRow(screen, cmd.A, cmd.B)
	case RotateColumn:
		runRotateColumn(screen, cmd.A, cmd.B)
	}
}

func runRect(screen *[][]bool, width, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			(*screen)[i][j] = true
		}
	}
}

func runRotateRow(screen *[][]bool, row, amount int) {
	width := len((*screen)[0])
	correctedAmount := amount % width
	for i := 0; i < correctedAmount; i++ {
		tmpPrev := (*screen)[row][0]
		var tmpNext bool
		for j := 0; j < width; j++ {
			tmpNext = (*screen)[row][j]
			(*screen)[row][j] = tmpPrev
			tmpPrev = tmpNext
		}
		(*screen)[row][0] = tmpPrev
	}
}

func runRotateColumn(screen *[][]bool, column, amount int) {
	heigth := len(*screen)
	correctedAmount := amount % heigth
	for i := 0; i < correctedAmount; i++ {
		tmpPrev := (*screen)[0][column]
		var tmpNext bool
		for j := 0; j < heigth; j++ {
			tmpNext = (*screen)[j][column]
			(*screen)[j][column] = tmpPrev
			tmpPrev = tmpNext
		}
		(*screen)[0][column] = tmpPrev
	}
}

func printScreen(screen [][]bool) {
	for _, row := range screen {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countLitPixels(screen [][]bool) int {
	count := 0
	for _, row := range screen {
		for _, col := range row {
			if col {
				count++
			}
		}
	}
	return count
}

func part2(data []Command, heigth, width int) {
	screen := make([][]bool, heigth)
	for i := range screen {
		screen[i] = make([]bool, width)
	}

	for _, cmd := range data {
		runCommand(&screen, cmd)
	}

	printScreen(screen)
}
