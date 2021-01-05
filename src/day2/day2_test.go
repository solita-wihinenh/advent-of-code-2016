package main

import (
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	data := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	got := part1(data)
	want := "1985"
	if got != want {
		t.Errorf("part1sample1 got = %q; want %q", got, want)
	}
}

func TestPart2Sample1(t *testing.T) {
	data := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	got := part2(data)
	want := "5DB3"
	if got != want {
		t.Errorf("part2sample1 got = %q; want %q", got, want)
	}
}
