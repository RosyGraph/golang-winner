package xbonacci

func Xbonacci(num, index int) int {
	a := 1
	b := 1

	for i := 0; i < index-2; i++ {
		a, b = a+b, a
	}
	return a
}
