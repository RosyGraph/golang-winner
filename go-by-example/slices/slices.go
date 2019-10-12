// Slices are a key data type in Go, giving a more powerful
// interface to sequences than arrays.
package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Printf("emp: %v\n", s)

	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Printf("set: %v\n", s)
	fmt.Printf("set: %v\n", s[2])

	fmt.Printf("len: %v\n", len(s))

	s = append(s, "e")
	s = append(s, "f", "g")
	fmt.Printf("apd: %v\n", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Printf("cp:  %v\n", c)

	l := s[2:5]
	fmt.Printf("sl1: %v\n", l)

	l = s[:5]
	fmt.Printf("sl2: %v\n", l)

	l = s[2:]
	fmt.Printf("sl3: %v\n", l)

	t := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Printf("dcl: %v\n", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Printf("2d:  %v\n", twoD)
}
