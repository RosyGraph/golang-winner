package sherlocksquares

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	lo   int32
	hi   int32
	want int32
}

func casesFromFile(t *testing.T) []testCase {
	t.Helper()

	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}

	cases := make([]testCase, 100)

	f, err := os.Open("testcases/testcases.txt")
	check(err)

	defer f.Close()

	s := bufio.NewScanner(f)
	for i := 0; s.Scan(); i++ {
		log.Println(s.Text())
		fields := strings.Fields(s.Text())

		parseInt := func(str string) int32 {
			n, err := strconv.Atoi(str)
			check(err)
			return int32(n)
		}

		cases[i] = testCase{
			parseInt(fields[0]),
			parseInt(fields[1]),
			parseInt(fields[2]),
		}
	}
	return cases
}

func TestSquares(t *testing.T) {
	tc := casesFromFile(t)
	/* tc := []struct {
		name string
		a    int32
		b    int32
		want int32
	}{
		{name: "test case 0", a: 3, b: 9, want: 2},
		{name: "test case 1", a: 17, b: 24, want: 0},
		{name: "test case 2", a: 4, b: 4, want: 1},
	} */

	for i, c := range tc {
		t.Run("test case 0", func(t *testing.T) {
			got := squares(c.lo, c.hi)

			if got != c.want {
				t.Errorf("test %d: got %d want %d", i, got, c.want)
			}
		})

	}
}
