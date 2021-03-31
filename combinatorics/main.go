package main

import (
	"errors"
	"fmt"
)

func fact(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("invalid n")
	}
	if n < 2 {
		return 1, nil
	}
	out, err := fact(n - 1)
	return n * out, err
}

func choose(n, k int) (int, error) {
	if n < 0 || k < 0 {
		return -1, errors.New("invalid n or k")
	}
	if k > n {
		return 0, nil
	}

	num, err := fact(n)
	if err != nil {
		return -1, err
	}

	d1, _ := fact(k)
	d2, _ := fact(n - k)
	denom := d1 * d2
	return num / denom, nil
}

func main() {
	for n := 0; n < 10; n++ {
		result, err := choose(n, 2)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(result)
		}
	}
}
