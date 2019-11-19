package main

import "testing"

func TestClosedIsland(t *testing.T) {
	tc := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "leetcode case 1",
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			want: 2,
		},
		{
			name: "leetcode case 2",
			grid: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			want: 1,
		},
		{
			name: "leetcode case 3",
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			want: 2,
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := closedIsland(c.grid)
			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}
