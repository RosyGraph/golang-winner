package main

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		n    int
		want int
	}{
		{name: "element in first half", arr: []int{0, 2, 4, 6, 8, 10}, n: 6, want: 3},
		{name: "element excluded", arr: []int{0, 2, 4, 6, 8, 10}, n: 5, want: -1},
		{name: "element in second half", arr: []int{0, 2, 4, 6, 8, 10}, n: 10, want: 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := BinarySearch(c.arr, c.n)
			if got != c.want {
				t.Errorf("got %v want %v given %v", got, c.want, c.arr)
			}
		})
	}
}
