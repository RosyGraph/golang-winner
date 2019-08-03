package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numberSlices ...[]int) (sums []int) {
	for _, numberSlice := range numberSlices {
		sums = append(sums, Sum(numberSlice))
	}

	return sums
}
