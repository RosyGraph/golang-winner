package main

import "testing"

/**
 * Return a copy of the input string with each 'a' character
 * replaced by '*'.
 *
 * @param string the string on which to perform the replacement
 * @return the string with each 'a' character replaced by '*'
 */
func TestHelloWorld(t *testing.T) {
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
	got := CensorA("All the ducks are swimming in the water")
	want := "*ll the ducks *re swimming in the w*ter"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
