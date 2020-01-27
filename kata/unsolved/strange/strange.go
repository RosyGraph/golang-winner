package kata

// inverse function
/* func strangeGrid(n int) []int {
	coords := []int{n / 10 * 2, n % 10 / 2}
	if n%2 != 0 {
		coords[0]++
	}
	return coords
} */

// FIXME: resolve test failures
func strangeGrid(row, col int32) int32 {
	// binary step arithmetic sequence
	r := uint64(row)
	c := uint64(col)
	var firstDigits uint64 = (r - 1) / 2 * 10
	var lastDigit uint64 = 2 * (c - 1)

	if row%2 == 0 { // if row is even incr
		lastDigit++
	}

	return int32(firstDigits + lastDigit)
}
