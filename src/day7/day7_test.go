package main

import (
	"testing"
)

func TestSupportTLS(t *testing.T) {
	cases := []struct {
		in   Ipv7
		want bool
	}{
		{Ipv7{[]string{"abba", "qrst"}, []string{"mnop"}}, true},
		{Ipv7{[]string{"abcd", "xyyx"}, []string{"bddb"}}, false},
		{Ipv7{[]string{"aaaa", "tyui"}, []string{"qwer"}}, false},
		{Ipv7{[]string{"ioxxoj", "zxcvbn"}, []string{"asdfgh"}}, true},
	}

	for _, c := range cases {
		got := c.in.SupportsTLS()
		if got != c.want {
			t.Errorf("%v got = %v; want %v", c.in, got, c.want)
		}
	}
}

func TestSupportsSSL(t *testing.T) {
	cases := []struct {
		in   Ipv7
		want bool
	}{
		{Ipv7{[]string{"aba", "xyz"}, []string{"bab"}}, true},
		{Ipv7{[]string{"xyx", "xyx"}, []string{"xyx"}}, false},
		{Ipv7{[]string{"aaa", "kek"}, []string{"eke"}}, true},
		{Ipv7{[]string{"zazbz", "cdb"}, []string{"bzb"}}, true},
	}

	for _, c := range cases {
		got := c.in.SupportsSSL()
		if got != c.want {
			t.Errorf("%v got = %v; want %v", c.in, got, c.want)
		}
	}
}
