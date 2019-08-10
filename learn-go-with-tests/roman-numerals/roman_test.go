package main

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		integer int
		roman   string
	}{
		{integer: 1, roman: "I"},
		{integer: 2, roman: "II"},
		{integer: 3, roman: "III"},
		{integer: 4, roman: "IV"},
		{integer: 5, roman: "V"},
		{integer: 6, roman: "VI"},
		{integer: 7, roman: "VII"},
		{integer: 8, roman: "VIII"},
		{integer: 9, roman: "IX"},
		{integer: 10, roman: "X"},
	}

	for _, tt := range cases {
		got := IntToRoman(tt.integer)
		if got != tt.roman {
			t.Errorf("got %s want %s given %d", got, tt.roman, tt.integer)
		}
	}
}
