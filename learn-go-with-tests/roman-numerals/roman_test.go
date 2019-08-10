package main

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		integer int
		roman   string
	}{
		{integer: 2, roman: "II"},
		{integer: 4, roman: "IV"},
		{integer: 5, roman: "V"},
		{integer: 7, roman: "VII"},
		{integer: 9, roman: "IX"},
		{integer: 10, roman: "X"},
		{integer: 11, roman: "XI"},
		{integer: 14, roman: "XIV"},
		{integer: 17, roman: "XVII"},
	}

	for _, tt := range cases {
		got := IntToRoman(tt.integer)
		if got != tt.roman {
			t.Errorf("got %s want %s given %d", got, tt.roman, tt.integer)
		}
	}
}
