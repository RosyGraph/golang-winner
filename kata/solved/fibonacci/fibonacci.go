package kata_fibonacci

type Fibonacci struct{}

func (f Fibonacci) Fib(n int) (res int) {
	if n == 1 || n == 2 {
		res = 1
	} else if n > 2 {
		res = f.Fib(n-1) + f.Fib(n-2)
	}
	return
}
