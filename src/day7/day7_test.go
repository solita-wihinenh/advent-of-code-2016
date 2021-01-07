package main

import (
	"testing"
)

func TestSupportTLSSample1(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"abba", "qrst"},
		Hypernets: []string{"mnop"},
	}
	got := data.SupportsTLS()
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportTLSSample2(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"abcd", "xyyx"},
		Hypernets: []string{"bddb"},
	}
	got := data.SupportsTLS()
	want := false
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportTLSSample3(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"aaaa", "tyui"},
		Hypernets: []string{"qwer"},
	}
	got := data.SupportsTLS()
	want := false
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportsTLSSample4(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"ioxxoj", "zxcvbn"},
		Hypernets: []string{"asdfgh"},
	}
	got := data.SupportsTLS()
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportsSSLSample1(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"aba", "xyz"},
		Hypernets: []string{"bab"},
	}
	got := data.SupportsSSL()
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportsSSLSample2(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"xyx", "xyx"},
		Hypernets: []string{"xyx"},
	}
	got := data.SupportsSSL()
	want := false
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportsSSLSample3(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"aaa", "kek"},
		Hypernets: []string{"eke"},
	}
	got := data.SupportsSSL()
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}

func TestSupportsSSLSample4(t *testing.T) {
	data := Ipv7{
		Sequences: []string{"zazbz", "cdb"},
		Hypernets: []string{"bzb"},
	}
	got := data.SupportsSSL()
	want := true
	if got != want {
		t.Errorf("%v got = %v; want %v", data, got, want)
	}
}
