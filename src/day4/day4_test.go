package main

import (
	"testing"
)

func TestIsRealRoomSample1(t *testing.T) {
	data := room{[]string{"aaaaa", "bbb", "z", "y", "x"}, 123, "abxyz"}
	got := isRealRoom(data)
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestIsRealRoomSample2(t *testing.T) {
	data := room{[]string{"a", "b", "c", "d", "e", "f", "g", "h"}, 987, "abcde"}
	got := isRealRoom(data)
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestIsRealRoomSample3(t *testing.T) {
	data := room{[]string{"not", "a", "real", "room"}, 404, "oarel"}
	got := isRealRoom(data)
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestIsRealRoomSample4(t *testing.T) {
	data := room{[]string{"totally", "real", "room"}, 200, "decoy"}
	got := isRealRoom(data)
	want := false
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}
