package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var answer chan int
var outputs map[int]int

func main() {
	fmt.Println("Day 10 - Start")
	data := loadData("day10part1.txt")
	answer = make(chan int)
	outputs = make(map[int]int)
	go executeInstructions(data)
	fmt.Println("Part 1 -", <-answer)
	fmt.Println("Part 2 -", outputs[0]*outputs[1]*outputs[2])
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

type bot struct {
	chips        []int
	number       int
	highIsOutput bool
	high         int
	lowIsOutput  bool
	low          int
}

func contains(slice []int, value int) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}

func (b *bot) checkAnswer() {
	if contains(b.chips, 61) && contains(b.chips, 17) {
		answer <- b.number
	}
}

func (b *bot) ReceiveChip(chip int) error {
	if len(b.chips) >= 2 {
		return fmt.Errorf("Already holding 2 chips %v", b.number)
	}
	b.chips = append(b.chips, chip)
	if len(b.chips) == 2 {
		b.checkAnswer()
	}
	return nil
}

func (b *bot) GiveLowerChip() (chip int, e error) {
	if len(b.chips) == 0 {
		e = errors.New("Not holding any chips")
		return
	}
	if len(b.chips) > 2 {
		e = errors.New("Holding more than 2 chips")
		return
	}

	if len(b.chips) == 2 {
		if b.chips[0] < b.chips[1] {
			chip = b.chips[0]
			b.chips = b.chips[1:]
		} else {
			chip = b.chips[1]
			b.chips = b.chips[:1]
		}
	} else if len(b.chips) == 1 {
		chip = b.chips[0]
		b.chips = b.chips[1:]
	}
	return
}

type instruction struct {
	value     int
	targetBot int
}

type inputData struct {
	instructions []instruction
	bots         map[int]*bot
}

func loadData(filename string) inputData {
	var data inputData
	data.bots = make(map[int]*bot)

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "value"):
			instruction := parseInstruction(line)
			data.instructions = append(data.instructions, instruction)
		case strings.HasPrefix(line, "bot"):
			bot := parseBot(line)
			data.bots[bot.number] = &bot
		default:
			panic(fmt.Sprintf("Unknown line type %v", line))
		}
	}
	return data
}

func parseInstruction(line string) instruction {
	var inst instruction
	i := 0
	inst.value, i = nextInt(line, i)
	inst.targetBot, i = nextInt(line, i)
	return inst
}

func parseBot(line string) bot {
	var b bot
	i := 0
	b.number, i = nextInt(line, i)
	b.low, i = nextInt(line, i)
	if line[i-len(strconv.Itoa(b.low))-4:i-len(strconv.Itoa(b.low))-1] != "bot" {
		b.lowIsOutput = true
	} else {
		b.lowIsOutput = false
	}
	b.high, i = nextInt(line, i)
	if line[i-len(strconv.Itoa(b.high))-4:i-len(strconv.Itoa(b.high))-1] != "bot" {
		b.highIsOutput = true
	} else {
		b.highIsOutput = false
	}
	return b
}

func executeInstructions(data inputData) {
	for _, v := range data.instructions {
		giveChipToBot(data.bots, outputs, v.value, v.targetBot)
	}
}

func giveChipToBot(bots map[int]*bot, outputs map[int]int, chip int, targetBot int) {
	b := bots[targetBot]
	err := b.ReceiveChip(chip)
	check(err)
	if len(b.chips) == 2 {
		lowerChip, err := b.GiveLowerChip()
		check(err)
		if b.lowIsOutput {
			outputs[b.low] = lowerChip
		} else {
			giveChipToBot(bots, outputs, lowerChip, b.low)
		}

		higherChip, err := b.GiveLowerChip()
		check(err)
		if b.highIsOutput {
			outputs[b.low] = higherChip
		} else {
			giveChipToBot(bots, outputs, higherChip, b.high)
		}
	}
}
