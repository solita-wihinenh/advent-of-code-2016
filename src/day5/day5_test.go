package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	data := "abc"
	got := part1(data)
	want := "18f47a30"
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestPart2(t *testing.T) {
	data := "abc"
	got := part2(data)
	want := "05ace8e3"
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}
