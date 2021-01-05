package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Day 2 - Start")
	data := loadData("day2part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadData(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.Split(string(bytes), "\n")
}

func part1(data []string) string {
	keyMap := map[[2]int]string{
		{0, 0}: "7",
		{1, 0}: "8",
		{2, 0}: "9",
		{0, 1}: "4",
		{1, 1}: "5",
		{2, 1}: "6",
		{0, 2}: "1",
		{1, 2}: "2",
		{2, 2}: "3",
	}
	return decode(data, keyMap, [2]int{1, 1})
}

func decode(data []string, keyMap map[[2]int]string, startPos [2]int) string {
	code := ""
	pos := startPos
	for _, line := range data {
		pos = decodeInstructions(pos, line, keyMap)
		key, ok := keyMap[pos]
		if !ok {
			panic(fmt.Sprintf("Key for position %v doesn't exist!", pos))
		}
		code += key
	}
	return code
}

func decodeInstructions(pos [2]int, line string, keyMap map[[2]int]string) [2]int {
	newPos := pos
	for _, v := range line {
		newPos = decodeInstruction(newPos, v, keyMap)
	}
	return newPos
}

func decodeInstruction(pos [2]int, instruction rune, keyMap map[[2]int]string) [2]int {
	nextPos := pos
	switch instruction {
	case 'R':
		nextPos = [2]int{pos[0] + 1, pos[1]}
	case 'L':
		nextPos = [2]int{pos[0] - 1, pos[1]}
	case 'U':
		nextPos = [2]int{pos[0], pos[1] + 1}
	case 'D':
		nextPos = [2]int{pos[0], pos[1] - 1}
	}
	if _, ok := keyMap[nextPos]; ok {
		return nextPos
	}
	return pos
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func part2(data []string) string {
	keyMap := map[[2]int]string{
		{2, 0}: "D",
		{1, 1}: "A",
		{2, 1}: "B",
		{3, 1}: "C",
		{0, 2}: "5",
		{1, 2}: "6",
		{2, 2}: "7",
		{3, 2}: "8",
		{4, 2}: "9",
		{1, 3}: "2",
		{2, 3}: "3",
		{3, 3}: "4",
		{2, 4}: "1",
	}
	return decode(data, keyMap, [2]int{0, 2})
}
