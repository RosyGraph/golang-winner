package main

func Sum(numbers []int) int {
	ans := 0
	for _, v := range numbers {
		ans += v
	}
	return ans
}

func SumAll(numberTable ...[]int) (sums []int) {
	for _, row := range numberTable {
		sums = append(sums, Sum(row))
	}
	return
}

func SumAllTails(numberTable ...[]int) (tails []int) {
	for _, row := range numberTable {
		if len(row) == 0 {
			tails = append(tails, 0)
		} else {
			tails = append(tails, Sum(row[1:]))
		}
	}
	return
}
