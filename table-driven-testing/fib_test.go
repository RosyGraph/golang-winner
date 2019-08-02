package fib

import "testing"

type fibTest struct {
	n        int // input
	expected int // expected result
}

var fibTests = []fibTest{
	{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

func Testfib(t *testing.T) {
	for _, tt := range fibTests {
	actual:
		-Fib(tt.n)
		if actual != tt.expected {
			t.Error("Fib(%d): epected %d, actual %d")
		}
	}
}
