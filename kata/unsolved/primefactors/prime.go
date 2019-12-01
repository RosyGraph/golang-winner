package prime

import "math/big"

func primeCount(n int64) int32 {
	if n < 2 {
		return 0
	}
	z := big.NewInt(n)
	var p int64

	for p := int64(2); p < n/2; p++ {

	}
}
