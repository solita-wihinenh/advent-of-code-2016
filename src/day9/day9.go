package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day 9 - Start")
	data := loadData("day9part1.txt")
	fmt.Println("Part 1 -", part1(data))
	fmt.Println("Part 2 -", part2(data))
}

func nextInt(str string, i int) (int, int) {
	for ; i < len(str) && !unicode.IsDigit(rune(str[i])); i++ {
	}
	x := 0
	for ; i < len(str) && unicode.IsDigit(rune(str[i])); i++ {
		x = x*10 + int(str[i]) - '0'
	}
	return x, i
}

func loadData(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return string(bytes)
}

func part1(data string) int {
	return len(decompress(data, DecompressV1))
}

// Marker length tells how many characters to multiply and multiplier how many times
type Marker struct {
	length, multiplier int
}

// DecompressFunction signature
type DecompressFunction func(Marker, string) (string, string)

// DecompressV1 returns decompressed part of given string, doesnt decompress markers
func DecompressV1(m Marker, compressed string) (unhandled, decompressed string) {
	decompressionResult := compressed[:m.length]
	for j := 0; j < m.multiplier; j++ {
		decompressed += decompressionResult
	}
	unhandled = compressed[m.length:]
	return
}

// DecompressV2 returns decompressed part of given string, decompresses markers
func DecompressV2(m Marker, compressed string) (unhandled, decompressed string) {
	decompressionResult := decompress(compressed[:m.length], DecompressV2)
	for j := 0; j < m.multiplier; j++ {
		decompressed += decompressionResult
	}
	unhandled = compressed[m.length:]
	return
}

func readMarker(str string) (Marker, string) {
	var length int
	var multiplier int
	i := 0
	length, i = nextInt(str, 0)
	multiplier, i = nextInt(str, i)
	return Marker{length, multiplier}, str[i+1:]
}

func decompress(compressed string, decompressionFn DecompressFunction) string {
	decompressed := ""
	unhandled := compressed
	for {
		leftParens := strings.Index(unhandled, "(")
		// No parens, read all
		if leftParens == -1 {
			decompressed += unhandled
			break
		}

		// Read up to parens
		decompressed += unhandled[:leftParens]
		unhandled = unhandled[leftParens:]

		var marker Marker
		marker, unhandled = readMarker(unhandled)
		var decompressedPart string
		unhandled, decompressedPart = decompressionFn(marker, unhandled)
		decompressed += decompressedPart
	}
	return decompressed
}

func part2(data string) int {
	return len(decompress(data, DecompressV2))
}
