package main

import (
	"reflect"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int, passed []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, passed)
		}
	}

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
