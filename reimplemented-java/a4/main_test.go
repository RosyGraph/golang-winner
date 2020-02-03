package a4

import "testing"

func TestMinPostiveInt(t *testing.T) {
	// t.Fatal("not implemented")
	tc := []struct {
		s    string
		want int
	}{
		{s: "3 4 2 11 666", want: 2},
		{s: "1", want: 1},
		{s: "1 0", want: 1},
	}

	for _, c := range tc {
		t.Run(c.s, func(t *testing.T) {
			got := MinPositiveInt(c.s)
			sameInt(t, got, c.want)
		})
	}
	got := MinPositiveInt("3 4 2 11 666")
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirstWordAlphabetically(t *testing.T) {
	tc := []struct {
		s    string
		want string
	}{
		{s: "def ga abe", want: "abe"},
		{s: "a", want: "a"},
	}

	for _, c := range tc {
		t.Run(c.s, func(t *testing.T) {
			got := FirstWordAlphabetically(c.s)
			sameString(t, got, c.want)
		})
	}
}

func TestMinIntTwoStr(t *testing.T) {
	tc := []struct {
		s1, s2 string
		want   int
	}{
		{s1: "32 5", s2: "32 1 4", want: 1},
		{s1: "0 1", s2: "32 1 4", want: 0},
		{s1: "1000", s2: "1000", want: 1000},
		{s1: "1000 8", s2: "9 -1000", want: -1000},
	}

	for _, c := range tc {
		s := c.s1 + " " + c.s2
		t.Run(s, func(t *testing.T) {
			got := MinIntTwoStr(c.s1, c.s2)
			sameInt(t, got, c.want)
		})
	}
}

func TestCurveScores(t *testing.T) {
	tc := []struct {
		s    string
		want string
	}{
		{s: "70 3 10", want: "100 33 40"},
		{s: "3 10", want: "93 100"},
		{s: "0", want: "100"},
	}

	for _, c := range tc {
		t.Run(c.s, func(t *testing.T) {
			sameString(t, CurveScores(c.s), c.want)
		})
	}
}

func sameString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func sameInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
