package sequence

func NthArithmetic(n, n1, d int) int {
	return n1 + (n-1)*d
}

// g_n = g_1 * r^{n-1}
func NthGeometric(n, n1, r int) int {
	return n1 * pow(r, (n-1))
}

func pow(b, p int) int {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return b
	}
	return b * pow(b, p-1)
}

// TODO implement series formulae
