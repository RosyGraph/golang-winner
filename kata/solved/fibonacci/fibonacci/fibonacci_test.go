package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	cases := []struct {
		name  string
		param int
		want  int
	}{
		{name: "Fib(0)", param: 0, want: 0},
		{name: "Fib(1)", param: 1, want: 1},
		{name: "Fib(2)", param: 2, want: 1},
		{name: "Fib(3)", param: 3, want: 2},
		{name: "Fib(4)", param: 4, want: 3},
		{name: "Fib(5)", param: 5, want: 5},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var fib Fibonacci
			got := fib.Fib(test.param)
			assertEquals(t, got, test.want)
		})
	}
}

func assertEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
