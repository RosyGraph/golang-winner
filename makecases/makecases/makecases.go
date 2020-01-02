package makecases

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	lo   int32
	hi   int32
	want int32
}

func NewCase(str string) []int32 {
	c := make([]int32, 3)
	parseField := func(s string) int32 {
		p, err := strconv.Atoi(s)
		if err != nil {
			os.Exit(1)
		}
		return int32(p)
	}

	f := strings.Fields(str)

	for i, v := range f {
		c[i] = parseField(v)
	}

	return c
}

func MakeCases() [][]int32 {
	readfile := func(s string) *bufio.Scanner {
		f, err := os.Open(s)
		if err != nil {
			os.Exit(1)
		}

		return bufio.NewScanner(f)
	}

	s1 := readfile("testcases/cases.txt")
	s2 := readfile("testcases/solutions.txt")

	i := 0
	var tc [][]int32
	for s1.Scan() {
		s2.Scan()
		s := fmt.Sprintf("%s %s", s1.Text(), s2.Text())

		tc[i] = NewCase(s)
		i++
	}
	return tc
}
