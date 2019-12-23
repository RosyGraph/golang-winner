package restaurant

func restaurant(n, m int32) int32 {
	var maxSize int32 = 1
	var lo int32 = min(n, m)

	for i := int32(2); i < lo+1; i++ {
		if divides(n, m, i) {
			maxSize = i
		}
	}
	return numSquares(n, m, maxSize)
}

func numSquares(n, m, d int32) int32 { return n / d * m / d }

func divides(n, m, i int32) bool { return n%i == 0 && m%i == 0 }

func min(n, m int32) int32 {
	if n > m {
		return n
	}
	return m
}
