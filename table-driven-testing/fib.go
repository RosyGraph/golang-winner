package fib

// Fib returns the nth number in the Fibonacci series.
// TODO: Implement threebonacci, etc
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-1)
}
