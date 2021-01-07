package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day 7 - Start")
	data := loadData("day7part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

// Ipv7 type
type Ipv7 struct {
	Sequences []string
	Hypernets []string
}

// SupportsTLS is true if ipv7 supports TLS
func (ipv7 Ipv7) SupportsTLS() bool {

	hasAbba := func(s string) bool {
		for i := 0; i < len(s)-3; i++ {
			if s[i] == s[i+1] {
				continue
			}
			reversed := string(s[i+1]) + string(s[i])
			if reversed == (s[i+2 : i+4]) {
				return true
			}
		}
		return false
	}

	sequenceAbba := false
	for _, sequence := range ipv7.Sequences {
		if hasAbba(sequence) {
			sequenceAbba = true
			break
		}
	}

	for _, hypernet := range ipv7.Hypernets {
		if hasAbba(hypernet) {
			return false
		}
	}

	return sequenceAbba
}

// SupportsSSL is true if ipv7 supports SSL
func (ipv7 Ipv7) SupportsSSL() bool {

	findAbas := func(s string) map[string]bool {
		abas := make(map[string]bool)
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+1] || s[i] != s[i+2] {
				continue
			}
			abas[s[i:i+3]] = true
		}
		return abas
	}

	abas := make(map[string]bool)
	for _, sequence := range ipv7.Sequences {
		seqAbas := findAbas(sequence)
		for aba := range seqAbas {
			abas[aba] = true
		}
	}

	for _, hypernet := range ipv7.Hypernets {
		for aba := range abas {
			reversed := string(aba[1]) + string(aba[0]) + string(aba[1])
			if strings.Contains(hypernet, reversed) {
				return true
			}
		}
	}

	return false
}

func loadData(filename string) []Ipv7 {
	var data []Ipv7

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, parseIpv7(scanner.Text()))
	}
	check(err)
	return data
}

func parseIpv7(str string) Ipv7 {
	var ipv7 Ipv7
	for len(str) > 0 {
		if str[0] == '[' {
			substr := str[1:strings.IndexByte(str, ']')]
			ipv7.Hypernets = append(ipv7.Hypernets, substr)
			str = str[len(substr)+2:]
		} else {
			bracketPos := strings.IndexByte(str, '[')
			var substr string
			if bracketPos != -1 {
				substr = str[:bracketPos]
			} else {
				substr = str
			}
			ipv7.Sequences = append(ipv7.Sequences, substr)
			str = str[len(substr):]
		}
	}
	return ipv7
}

func part1(data []Ipv7) int {
	count := 0
	for _, ipv7 := range data {
		if ipv7.SupportsTLS() {
			count++
		}
	}
	return count
}

func part2(data []Ipv7) int {
	count := 0
	for _, ipv7 := range data {
		if ipv7.SupportsSSL() {
			count++
		}
	}
	return count
}
