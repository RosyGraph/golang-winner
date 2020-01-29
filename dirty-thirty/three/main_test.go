package main

import "testing"

func TestWeird(t *testing.T) {
	tc := []struct {
		name string
		n    int32
		want string
	}{
		{name: "3", n: 3, want: "Weird"},
		{name: "4", n: 4, want: "Not Weird"},
		{name: "24", n: 24, want: "Not Weird"},
		{name: "8", n: 8, want: "Weird"},
	}

	for _, c := range tc {
		got := weird(c.n)

		if got != c.want {
			t.Errorf("given '%d' got '%s' want '%s'", c.n, got, c.want)
		}
	}
}
