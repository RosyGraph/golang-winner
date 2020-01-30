package main

import "testing"

func TestPerson(t *testing.T) {
	var tc = []struct {
		name string
		age  int
		want int
	}{
		{name: "10", age: 10, want: 10},
		{name: "99", age: 99, want: 99},
		{name: "0", age: 0, want: 0},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			p := person{age: c.age}

			got := p.age

			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}
