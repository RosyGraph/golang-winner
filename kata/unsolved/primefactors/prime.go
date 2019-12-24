package prime

import (
	"math"
	"math/big"
)

func primeCount(n int64) int32 {
	if n < 2 {
		return 0
	}
	if isPrime(n) {
		return 1
	}

	var count int32

	if n%2 == 0 {
		count++
	}

	for n%2 == 0 {
		n /= 2
	}

	nroot := int64(math.Sqrt(float64(n)))
	for i := int64(3); i < nroot; {
		if n%i == 0 && isPrime(i) {
			count++
			n /= i
		} else {
			i += 2
		}
	}

	if n > 2 && isPrime(n) {
		count++
	}

	return count
}

func isPrime(n int64) bool {
	x := new(big.Int)
	x.SetInt64(n)

	return x.ProbablyPrime(1)
}
