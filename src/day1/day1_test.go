package main

import (
	"fmt"
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	data := []string{"R2", "L3"}
	got := part1(data)
	want := 5
	if got != want {
		t.Errorf("part1sample1 got = %d; want %d", got, want)
	}
}

func TestPart1Sample2(t *testing.T) {
	data := []string{"R2", "R2", "R2"}
	got := part1(data)
	want := 2
	if got != want {
		t.Errorf("part1sample2 got = %d; want %d", got, want)
	}
}

func TestPart1Sample3(t *testing.T) {
	data := []string{"R5", "L5", "R5", "R3"}
	got := part1(data)
	want := 12
	if got != want {
		t.Errorf("part1sample3 got = %d; want %d", got, want)
	}
}

func TestPart2Sample1(t *testing.T) {
	fmt.Println("test2")
	data := []string{"R8", "R4", "R4", "R8"}
	got := part2(data)
	want := 4
	if got != want {
		t.Errorf("part2sample1 got = %d; want %d", got, want)
	}
}
