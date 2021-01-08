package main

import "testing"

func TestDecompressV1(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"ADVENT", "ADVENT"},
		{"A(1x5)BC", "ABBBBBC"},
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		{"(6x1)(1x3)A", "(1x3)A"},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
	}

	for _, c := range cases {
		got := decompress(c.in, DecompressV1)
		if got != c.want {
			t.Errorf("%v got = %v; want %v", c.in, got, c.want)
		}
	}
}

func TestDecompressV2(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"X(8x2)(3x3)ABCY", "XABCABCABCABCABCABCY"},
		//{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		//{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for _, c := range cases {
		got := decompress(c.in, DecompressV2)
		if got != c.want {
			t.Errorf("%v got = %v; want %v", c.in, got, c.want)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for _, c := range cases {
		got := part2(c.in)
		if got != c.want {
			t.Errorf("%v got = %v; want %v", c.in, got, c.want)
		}
	}
}
