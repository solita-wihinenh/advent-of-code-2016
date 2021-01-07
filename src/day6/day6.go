package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day 6 - Start")
	data := loadData("day6part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

func loadData(filename string) []string {
	var data []string

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	check(err)
	return data
}

type keyCount struct {
	key   rune
	count int
}

func getSortedKeyCounts(dict map[rune]int) []keyCount {
	var keyCounts []keyCount
	for char, count := range dict {
		keyCounts = append(keyCounts, keyCount{char, count})
	}

	sort.Slice(keyCounts, func(first, second int) bool {
		if keyCounts[first].count > keyCounts[second].count {
			return true
		} else if keyCounts[first].count < keyCounts[second].count {
			return false
		}
		return keyCounts[first].key < keyCounts[second].key
	})
	return keyCounts
}

func part1(data []string) string {
	dicts := populateKeyCountDictionaries(data)

	message := ""
	for _, dict := range dicts {
		keyCounts := getSortedKeyCounts(dict)
		message += string(keyCounts[0].key)
	}

	return message
}

func populateKeyCountDictionaries(data []string) []map[rune]int {
	var dicts []map[rune]int

	for i := 0; i < len(data[0]); i++ {
		dicts = append(dicts, make(map[rune]int))
	}

	for _, str := range data {
		for i, char := range str {
			if _, ok := dicts[i][char]; !ok {
				dicts[i][char] = 1
			} else {
				dicts[i][char]++
			}
		}
	}
	return dicts
}

func part2(data []string) string {
	dicts := populateKeyCountDictionaries(data)

	message := ""
	for _, dict := range dicts {
		keyCounts := getSortedKeyCounts(dict)
		message += string(keyCounts[len(keyCounts)-1].key)
	}

	return message
}
