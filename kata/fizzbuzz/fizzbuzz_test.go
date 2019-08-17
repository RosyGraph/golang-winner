package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		name string
		arg  int
		want string
	}{
		{name: "1 = 1", arg: 1, want: "1"},
		{name: "2 = 2", arg: 2, want: "2"},
		{name: "3 = Fizz", arg: 3, want: "Fizz"},
		{name: "4 = 4", arg: 4, want: "4"},
		{name: "5 = Buzz", arg: 5, want: "Buzz"},
		{name: "6 = Fizz", arg: 6, want: "Fizz"},
		{name: "7 = 7", arg: 7, want: "7"},
		{name: "8 = 8", arg: 8, want: "8"},
		{name: "9 = Fizz", arg: 9, want: "Fizz"},
		{name: "10 = Buzz", arg: 10, want: "Buzz"},
		{name: "11 = 11", arg: 11, want: "11"},
		{name: "12 = Fizz", arg: 12, want: "Fizz"},
		{name: "15 = FizzBuzz", arg: 15, want: "FizzBuzz"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assertEqual(t, test.arg, test.want)
		})
	}
}

func assertEqual(t *testing.T, n int, want string) {
	t.Helper()
	got := FizzBuzz(n)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
