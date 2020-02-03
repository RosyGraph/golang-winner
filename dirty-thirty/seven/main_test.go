package main

import "testing"

func TestHelloWorld(t *testing.T) {
	tc := []struct {
		name string
		arr  []int32
		want string
	}{
		{name: "1 4 3 2", arr: []int32{1, 4, 3, 2}, want: "2 3 4 1"},
		{name: "2 2 3", arr: []int32{2, 2, 3}, want: "3 2 2"},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := reversed(c.arr)
			if got != c.want {
				t.Errorf("got '%s' want '%s'", got, c.want)
			}
		})
	}
}
