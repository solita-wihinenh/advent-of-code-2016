package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day 5 - Start")
	fmt.Println("Part 1 -", part1("reyedfim"))
	fmt.Println("Part 2 -", part2("reyedfim"))
}

func part1(doorID string) string {
	password := ""
	for i := 0; i <= math.MaxInt64; i++ {
		secret := doorID + strconv.Itoa(i)
		hash := generateMD5(secret)
		hexString := fmt.Sprintf("%x\n", hash)
		if hexString[0:5] == "00000" {
			password += hexString[5:6]
		}
		if len(password) == 8 {
			break
		}
	}
	return password
}

func generateMD5(input string) []byte {
	hash := md5.New()
	io.WriteString(hash, input)
	return hash.Sum(nil)
}

func part2(doorID string) string {
	password := make([]rune, 8)
	decodeCount := 0
	for i := 0; i <= math.MaxInt64 && decodeCount < 8; i++ {
		secret := doorID + strconv.Itoa(i)
		hash := generateMD5(secret)
		hexString := fmt.Sprintf("%x\n", hash)
		if hexString[0:5] == "00000" {
			pos, err := strconv.Atoi(hexString[5:6])
			if err != nil || pos < 0 || pos > 7 || password[pos] != 0 {
				continue
			}
			password[pos] = []rune(hexString[6:7])[0]
			decodeCount++
		}
	}
	return string(password)
}
