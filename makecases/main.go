package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	f1, err := os.Open("testcases/cases.txt")
	check(err)

	f2, err := os.Open("testcases/solutions.txt")

	defer f1.Close()
	defer f2.Close()

	s1 := bufio.NewScanner(f1)
	s2 := bufio.NewScanner(f2)

	for s1.Scan() {
		s2.Scan()
		fmt.Printf("%s %s\n", s1.Text(), s2.Text())
	}

}
