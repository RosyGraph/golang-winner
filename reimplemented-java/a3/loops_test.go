package a3

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

func TestTriangleString(t *testing.T) {
	tc := []struct {
		n    int
		want string
	}{
		{
			n:    3,
			want: "*\n**\n***",
		},
		{
			n:    1,
			want: "*",
		},
	}

	for _, c := range tc {
		t.Run(c.want, func(t *testing.T) {
			got := TriangleString(c.n)

			if got != c.want {
				t.Errorf("TriangleString(%d):\ngot:\n%s\nwant:\n%s", c.n, got, c.want)
			}
		})
	}
}

func TestSafeColor(t *testing.T) {
	tc := []struct {
		n    int
		want int
	}{
		{n: 288, want: 255},
		{n: -8, want: 0},
		{n: 20, want: 20},
	}

	for _, c := range tc {
		t.Run("SafeColor", func(t *testing.T) {
			got := SafeColor(c.n)

			if got != c.want {
				t.Errorf("SafeColor(%d): got %d want %d", c.n, got, c.want)
			}
		})
	}
}
