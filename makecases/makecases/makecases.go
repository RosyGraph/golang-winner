package makecases

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	Lo   int32
	Hi   int32
	Want int32
}

func NewCase(str string) (c Case) {
	parseField := func(s string) int32 {
		p, err := strconv.Atoi(s)
		if err != nil {
			os.Exit(1)
		}
		return int32(p)
	}

	f := strings.Fields(str)

	c.Lo = parseField(f[0])
	c.Hi = parseField(f[1])
	c.Want = parseField(f[2])

	return c
}

func MakeCases() []Case {
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
	var tc []Case
	for s1.Scan() {
		s2.Scan()
		s := fmt.Sprintf("%s %s", s1.Text(), s2.Text())

		tc[i] = NewCase(s)
		i++
	}
	return tc
}
