package main

import (
	//"fmt"
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	data := [][]int{{5, 10, 25}}
	got := part1(data)
	want := 0
	if got != want {
		t.Errorf("part1sample2 got = %q; want %q", got, want)
	}
}

func TestPart1Sample2(t *testing.T) {
	data := [][]int{{5, 25, 10}}
	got := part1(data)
	want := 0
	if got != want {
		t.Errorf("part1sample2 got = %q; want %q", got, want)
	}
}

func TestPart1Sample3(t *testing.T) {
	data := [][]int{{1, 2, 2}}
	got := part1(data)
	want := 1
	if got != want {
		t.Errorf("part1sample3 got = %q; want %q", got, want)
	}
}

func TestPart1Sample4(t *testing.T) {
	data := [][]int{{2, 2, 1}}
	got := part1(data)
	want := 1
	if got != want {
		t.Errorf("part1sample4 got = %q; want %q", got, want)
	}
}

func TestPart2Sample1(t *testing.T) {
	data := [][]int{{1, 2, 1}, {1, 2, 1}, {1, 4, 1}, {1, 2, 1}, {1, 2, 1}, {1, 4, 1}}
	got := part2(data)
	want := 4
	if got != want {
		t.Errorf("part2sample1 got = %q; want %q", got, want)
	}
}
