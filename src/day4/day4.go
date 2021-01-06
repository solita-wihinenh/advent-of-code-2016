package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 4 - Start")
	data := loadData("day4part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type room struct {
	encryptedName []string
	sectorID      int
	checksum      string
}

func loadData(filename string) []room {
	data := []room{}

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lineTokens := strings.Split(scanner.Text(), "-")
		encryptedName := lineTokens[:len(lineTokens)-1]
		lastToken := lineTokens[len(lineTokens)-1]
		sectorID, err := strconv.Atoi(lastToken[:strings.IndexByte(lastToken, '[')])
		check(err)
		checksum := lastToken[strings.IndexByte(lastToken, '[')+1 : len(lastToken)-1]
		data = append(data, room{encryptedName, sectorID, checksum})
	}
	check(err)
	return data
}

func part1(data []room) int {
	realRooms := findRealRooms(data)
	sectorTotal := 0
	for _, r := range realRooms {
		sectorTotal += r.sectorID
	}
	return sectorTotal
}

type keyCount struct {
	key   rune
	count int
}

func isRealRoom(r room) bool {
	dict := map[rune]int{}
	for _, key := range r.encryptedName {
		for _, char := range key {
			if _, ok := dict[char]; ok {
				dict[char]++
			} else {
				dict[char] = 1
			}
		}
	}

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

	for i, checksumChar := range r.checksum {
		if keyCounts[i].key != checksumChar {
			return false
		}
	}

	return true
}

func findRealRooms(rooms []room) []room {
	realRooms := []room{}
	for _, r := range rooms {
		if isRealRoom(r) {
			realRooms = append(realRooms, r)
		}
	}
	return realRooms
}

func part2(data []room) int {
	realRooms := findRealRooms(data)
	for _, r := range realRooms {
		decryptedName := decryptRoomName(r)
		//fmt.Printf("%v: %d\n", decryptedName, r.sectorID)
		if decryptedName == "northpole object storage" {
			return r.sectorID
		}
	}
	panic("No room with name: northpole object storage")
}

func decryptRoomName(r room) string {
	decryptedName := ""
	for i, word := range r.encryptedName {
		for _, char := range word {
			decryptedChar := int(char) + r.sectorID
			for ; decryptedChar > 122; decryptedChar -= 26 {
			}
			decryptedName += fmt.Sprintf("%c", decryptedChar)
		}
		if i != len(r.encryptedName)-1 {
			decryptedName += " "
		}
	}
	return decryptedName
}
