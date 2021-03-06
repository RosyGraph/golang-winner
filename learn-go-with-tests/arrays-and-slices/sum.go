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

	return
}

func SumAllTails(numberSlices ...[]int) (sumOfTails []int) {
	for _, numberSlice := range numberSlices {
		if len(numberSlice) < 1 {
			sumOfTails = append(sumOfTails, 0)
		} else {
			tail := numberSlice[1:]
			sumOfTails = append(sumOfTails, Sum(tail))
		}
	}

	return
}
