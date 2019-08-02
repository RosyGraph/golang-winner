package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertEquals := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if want != got {
			t.Errorf("Given %v, got %d want %d.", numbers, got, want)
		}
	}

	t.Run("Slice of length 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertEquals(t, got, want, numbers)
	})

	t.Run("Slice of length 3", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertEquals(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("Make the sums of 3 slices' tails.", func(t *testing.T) {
		got := SumAllTails([]int{3, 6}, []int{0, 2}, []int{8, 1})
		want := []int{6, 2, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v.", got, want)
		}
	})

	t.Run("Handle empty slices.", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v.", got, want)
		}
	})
}
