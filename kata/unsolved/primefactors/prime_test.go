package prime

import "testing"

type TestCase struct {
	name string
	n    int64
	want int32
}

func TestIsPrime(t *testing.T) {
	tc := []struct {
		name string
		n    int64
		want bool
	}{
		{name: "13", n: 13, want: true},
		{name: "3", n: 3, want: true},
		{name: "6", n: 6, want: false},
		{name: "1", n: 1, want: false},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := isPrime(c.n)
			if got != c.want {
				t.Errorf("give %s: got %v want %v", c.name, got, c.want)
			}
		})
	}
}

func TestPrimeCount(t *testing.T) {
	tc := []TestCase{
		{name: "1", n: 1, want: 0},
		{name: "2", n: 2, want: 1},
		{name: "3", n: 3, want: 1},
		{name: "500", n: 500, want: 4},
		{name: "5000", n: 5000, want: 5},
		{name: "10000000000", n: 10000000000, want: 10},
	}

	check := func(t *testing.T, c TestCase) {
		got := primeCount(c.n)
		if got != c.want {
			t.Errorf("given %s: got %d want %d", c.name, got, c.want)
		}
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			check(t, c)
		})
	}
}
