package xbonacci

func Xbonacci(num, index int) int {
	if index <= num {
		return 1
	}

	arr := make([]int, index)
	for i := 0; i < num; i++ {
		arr[i] = 1
	}

	for i := num; i < index; i++ {
		for j := i - num; j < num; j++ {
			arr[i] += arr[j]
		}
	}
	return arr[index-1]
}

func sum(arr []int) (s int) {
	for _, n := range arr {
		s += n
	}

	return s
}
