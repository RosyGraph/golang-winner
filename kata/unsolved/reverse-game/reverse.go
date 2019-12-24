package kata

func ReverseGame(n, k int32) int32 {
	b := reverse(newBalls(n))
	for i := int32(0); i < n-1; i++ {
		b = append(b[:i], reverse(b[i:])...)
	}
	return b[k]
}

func newBalls(n int32) []int32 {
	res := make([]int32, n)
	for i := int32(0); i < n; i++ {
		res[i] = i
	}
	return res
}

func reverse(n []int32) []int32 {
	if len(n) < 2 {
		return n
	}

	return append(reverse(n[1:]), n[0])
}
