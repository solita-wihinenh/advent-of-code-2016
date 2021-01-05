package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1 - Start")
	data := loadData("day1part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

// Direction enum
type Direction int

// Direction enum value
const (
	North Direction = iota
	East            = iota
	South           = iota
	West            = iota
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadData(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.Split(string(bytes), ", ")
}

func manhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	xDist := int(math.Abs(float64(x1 - x2)))
	yDist := int(math.Abs(float64(y1 - y2)))
	return xDist + yDist
}

func part1(data []string) int {
	x := 0
	y := 0
	heading := North
	for _, v := range data {
		switch v[0] {
		case 'R':
			heading = turnRight(heading)
		case 'L':
			heading = turnLeft(heading)
		}
		amount, err := strconv.Atoi(v[1:])
		check(err)
		switch heading {
		case North:
			y += amount
		case South:
			y -= amount
		case East:
			x += amount
		case West:
			x -= amount
		}
	}
	return manhattanDistance(0, 0, x, y)
}

func turnLeft(currentHeading Direction) Direction {
	return (currentHeading + 3) % 4
}

func turnRight(currentHeading Direction) Direction {
	return (currentHeading + 1) % 4
}

func part2(data []string) int {
	heading := North
	x := 0
	y := 0
	visited := map[[2]int]bool{
		{0, 0}: true,
	}

	for _, v := range data {
		switch v[0] {
		case 'R':
			heading = turnRight(heading)
		case 'L':
			heading = turnLeft(heading)
		}
		amount, err := strconv.Atoi(v[1:])
		check(err)
		for i := 0; i < amount; i++ {
			switch heading {
			case North:
				y++
			case South:
				y--
			case East:
				x++
			case West:
				x--
			}
			if _, ok := visited[[2]int{x, y}]; ok {
				return manhattanDistance(0, 0, x, y)
			}
			visited[[2]int{x, y}] = true
		}
	}
	panic("No location was visited twice!")
}
