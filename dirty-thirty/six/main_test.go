package main

import "testing"

func TestHelloWorld(t *testing.T) {
	tc := []struct {
		name string
		s    string
		want string
	}{
		{name: "test case 0", s: "Hacker", want: "Hce akr"},
		{name: "test case 1", s: "Rank", want: "Rn ak"},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := separate(c.s)

			if got != c.want {
				t.Errorf("got '%s' want '%s'", got, c.want)
			}
		})
	}

}
