package main

import (
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	data := loadData("day6part1sample1.txt")
	got := part1(data)
	want := "easter"
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestPart2Sample1(t *testing.T) {
	data := loadData("day6part1sample1.txt")
	got := part2(data)
	want := "advent"
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}
