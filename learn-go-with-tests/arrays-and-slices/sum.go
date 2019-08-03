package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numberSlices ...[]int) []int {
	lengthOfSlices := len(numberSlices)
	sums := make([]int, lengthOfSlices)

	for i, numberSlice := range numberSlices {
		sums[i] = Sum(numberSlice)
	}

	return sums
}
