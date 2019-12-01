package xbonacci

func Xbonacci(num, index int) int {
	arr := new([]int)
	for i := 0; i < num; i++ {
		*arr = append(*arr, 1)
	}

	a := 1

	for i := 0; i < index-num; i++ {
	}
	return a
}

func sum(arr []int) (s int) {
	for _, n := range arr {
		s += n
	}

	return s
}
