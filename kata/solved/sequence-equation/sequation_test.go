package main

import (
	"reflect"
	"testing"
)

func TestPermutationEquation(t *testing.T) {
	tc := []struct {
		name string
		g    []int32
		want []int32
	}{
		{name: "input 0", g: []int32{4, 3, 5, 1, 2}, want: []int32{1, 3, 5, 4, 2}},
		{name: "input 1", g: []int32{2, 3, 1}, want: []int32{2, 3, 1}},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := permutationEquation(c.g)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}
