package main

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		in   []Command
		want int
	}{
		{[]Command{{Rect, 3, 2}, {RotateColumn, 1, 1}, {RotateRow, 0, 4}, {RotateColumn, 1, 1}}, 6},
		{[]Command{{Rect, 1, 3}}, 3},
		{[]Command{{Rect, 2, 3}}, 6},
		{[]Command{{Rect, 3, 2}, {RotateRow, 0, 5}}, 6},
	}

	for _, c := range cases {
		got := part1(c.in, 3, 7)
		if got != c.want {
			t.Errorf("%v got = %v; want %v", c.in, got, c.want)
		}
	}
}
