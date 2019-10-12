// hackerrank.com kata "Sherlock and Squares"
// https://www.hackerrank.com/challenges/sherlock-and-squares/

// Statement
////////////////////////////////////////////////////////////////////////////////

/*
 * Watson likes to challenge Sherlock's math ability. He will provide a starting
 * and ending value describing a range of integers. Sherlock must determine the
 * number of square integers within that range, inclusive of the endpoints.
 *     Note : A square integer is an integer which is the square of an integer,
 *     e.g. 1, 4, 9, 16, 25.
 * For example, the range is a = 24 and b = 49, inclusive. There are three square
 * integers in the range: 25, 36, 49.
 */

// Function Description
////////////////////////////////////////////////////////////////////////////////

/*
 * Complete the squares function in the editor below. It should return an integer
 * representing the number of square integers in the inclusive range from a to b.
 *     squares has the following parameters:
 *     a: an integer, the lower range boundary
 *     b: an integer, the upper range boundary
 */

// Input Format
////////////////////////////////////////////////////////////////////////////////

/*
 * The first line contains q, the number of test cases.
 * Each of the next q lines contains two space-separated integers denoting a and
 * b, the starting and ending integers in the ranges.
 */

// Constraints
////////////////////////////////////////////////////////////////////////////////

/*
 * 1 <= q <= 100
 * 1 <= a <= b <= 10^{9}
 */

// Output Format
////////////////////////////////////////////////////////////////////////////////

/*
 * For each test case, return an int32.
 */

package sherlocksquares

import (
	"errors"
	"math"
)

type HashsetMap map[float64]struct{}

type Hashset struct {
	set HashsetMap
}

func (h *Hashset) Add(value float64) error {
	if _, ok := h.set[value]; ok {
		return errors.New("cannot add duplicate values")
	}

	h.set[value] = struct{}{}
	return nil
}

func New(values ...float64) *Hashset {
	hashset := Hashset{make(HashsetMap)}

	for _, v := range values {
		hashset.set[v] = struct{}{}
	}

	return &hashset
}

func squares(a, b int32) int32 {
	var c int32

	h := New()
	for i := float64(0); i < 10000; i++ {
		h.Add(math.Pow(i, 2))
	}

	for a <= b {
		_, ok := h.set[float64(a)]
		if ok {
			c++
		}
		a++
	}
	return c
}
