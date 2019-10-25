package main

import "testing"

func TestHourglassSum(t *testing.T) {
	a := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	got := HourglassSum(a)
	want := int32(19)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
