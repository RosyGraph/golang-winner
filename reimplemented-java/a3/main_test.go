package main

import "testing"

func TestCensorA(t *testing.T) {
	tc := []struct {
		s    string
		want string
	}{
		{
			s:    "All the ducks are swimming in the water",
			want: "*ll the ducks *re swimming in the w*ter",
		},
		{
			s:    "aAaaa",
			want: "*****",
		},
		{
			s:    "",
			want: "",
		},
	}

	for _, c := range tc {
		t.Run(c.s, func(t *testing.T) {
			got := CensorA(c.s)
			if got != c.want {
				t.Errorf("CensorA(%s):\ngot '%s'\nwant '%s'",
					c.s, got, c.want)
			}
		})
	}
}

func TestContainsMoreEvens(t *testing.T) {
	tc := []struct {
		s    string
		want bool
	}{
		{
			s:    "1 1 1 2",
			want: false,
		},
		{
			s:    "2 1 1 2 9 8 8",
			want: true,
		},
		{
			s:    "a2 1 1 2 9 8 8",
			want: false,
		},
		{
			s:    "",
			want: false,
		},
	}

	for _, c := range tc {
		t.Run(c.s, func(t *testing.T) {
			got := ContainsMoreEvens(c.s)
			if got != c.want {
				t.Errorf("ContainsMoreEvens(%s):\ngot '%v'\nwant '%v'", c.s, got, c.want)
			}
		})
	}
}
