package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3 - Start")
	data := loadData("day3part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadData(filename string) [][]int {
	var data [][]int
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		var triangle []int
		sides := strings.Split(line, " ")
		for _, side := range sides {
			side = strings.Replace(side, "\r", "", -1)
			if side == "" {
				continue
			}
			length, err := strconv.Atoi(side)
			check(err)
			triangle = append(triangle, length)
		}
		data = append(data, triangle)
	}
	return data
}

func part1(data [][]int) int {
	count := 0
	for _, row := range data {
		if isTriangle(row) {
			count++
		}
	}
	return count
}

func isTriangle(shape []int) bool {
	sorted := make([]int, len(shape))
	copy(sorted, shape)
	total := 0
	sort.Ints(sorted)
	for i, side := range sorted {
		if i == len(sorted)-1 && total <= side {
			return false
		}
		total += side
	}
	return true
}

func part2(data [][]int) int {
	count := 0
	for i := 0; i < len(data)-2; i += 3 {
		for j := 0; j < 3; j++ {
			triangle := []int{data[i][j], data[i+1][j], data[i+2][j]}
			if isTriangle(triangle) {
				count++
			}
		}
	}
	return count
}
