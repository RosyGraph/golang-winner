package kata

import (
	"reflect"
	"testing"
)

func TestReverseGame(t *testing.T) {
	tc := []struct {
		name string
		n, k int32
		want int32
	}{
		{name: "3 1", n: 3, k: 1, want: 2},
		{name: "5 2", n: 5, k: 2, want: 4},
		{name: "11 1", n: 11, k: 1, want: 3},
		{name: "3 1", n: 3, k: 1, want: 2},
		{name: "9 0", n: 9, k: 0, want: 1},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := ReverseGame(c.n, c.k)
			if got != c.want {
				t.Errorf("give %s: got %d want %d", c.name, got, c.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	got := reverse([]int32{0, 1, 2})
	want := []int32{2, 1, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNewBalls(t *testing.T) {
	tc := []struct {
		name string
		n    int32
		want []int32
	}{
		{name: "3", n: 3, want: []int32{0, 1, 2}},
		{name: "5", n: 5, want: []int32{0, 1, 2, 3, 4}},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := newBalls(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("given %s: got %v want %v", c.name, got, c.want)
			}
		})
	}
	got := newBalls(3)
	want := []int32{0, 1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
