package kata

func numJewelsInStones(j, s string) (res int) {
	m := make(map[rune]int)
	for _, r := range j {
		m[r] = 1
	}

	for _, r := range s {
		res += m[r]
	}

	return res
}
