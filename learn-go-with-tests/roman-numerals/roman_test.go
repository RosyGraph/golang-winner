package main

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		integer  int
		romanNum string
	}{
		{integer: 1, romanNum: "I"},
		{integer: 2, romanNum: "II"},
		{integer: 3, romanNum: "III"},
	}

	for _, tt := range cases {
		got := IntToRoman(tt.integer)
		if got != tt.romanNum {
			t.Errorf("got %s want %s given %d", got, tt.romanNum, tt.integer)
		}
	}
}
