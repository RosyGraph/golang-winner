package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	tc := []struct {
		name string
		n, m int
		want Matrix
	}{
		{name: "2x2", n: 2, m: 2, want: Matrix{{0, 0}, {0, 0}}},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := NewMatrix(c.n, c.m)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("case %s: got %v want %v", c.name, got, c.want)
			}
		})
	}
}
