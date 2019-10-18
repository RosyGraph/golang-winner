package main

import (
	"errors"
	"fmt"
	"math"

	sherlocksquares "github.com/RosyGraph/kata/unsolved/sherlocksquares/squares"
)

type HashsetMap map[float64]struct{}

type Hashset struct {
	set HashsetMap
}

func (h *Hashset) add(value float64) error {
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

func main() {
	h := sherlocksquares.New()
	for i := 0; i < 1000; i++ {
		h.Add(math.Pow(float64(i), 2))
	}
	fmt.Printf("%#v\n", h)
}
