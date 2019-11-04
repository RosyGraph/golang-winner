package main

import "testing"

func TesthourglassSum(t *testing.T) {
	tc := []struct {
		n    string
		a    [][]int32
		want int32
	}{
		{
			n: "positive numbers",
			a: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0}},
			want: 19,
		},
		{
			n: "positive numbers",
			a: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0}},
			want: 7,
		},
	}
	for _, c := range tc {
		t.Run(c.n, func(t *testing.T) {
			got := hourglassSum(c.a)
			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}

func TestSumGlass(t *testing.T) {
	a := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	got := sumGlass(a, 1, 1)
	want := int32(7)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
