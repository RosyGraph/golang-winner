package main

import (
	"reflect"
	"testing"
)

func TestFindPoint(t *testing.T) {
	tc := []struct {
		name           string
		px, py, qx, qy int32
		want           []int32
	}{
		{"pos slope", 1, 1, 2, 2, []int32{3, 3}},
		{"", 1, 1, 2, 2, []int32{3, 3}},
		{"", 4, 3, 5, 2, []int32{6, 1}},
		{"", 2, 4, 5, 6, []int32{8, 8}},
		{"", 1, 2, 2, 2, []int32{3, 2}},
		{"", 1, 1, 1, 1, []int32{1, 1}},
		{"", 1, 2, 2, 1, []int32{3, 0}},
		{"", 1, 8, 7, 8, []int32{13, 8}},
		{"", 9, 1, 1, 1, []int32{-7, 1}},
		{"", 8, 4, 3, 2, []int32{-2, 0}},
		{"", 7, 8, 9, 1, []int32{11, -6}},
	}

	for _, c := range tc {
		got := findPoint(c.px, c.py, c.qx, c.qy)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
