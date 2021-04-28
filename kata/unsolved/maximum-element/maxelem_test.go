package main

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tc := []struct {
		name    string
		queries []string
		want    []int32
	}{
		{name: "push one", queries: []string{"1 1"}, want: []int32{}},
		{name: "push two", queries: []string{"1 2"}, want: []int32{}},
		{name: "push 100", queries: []string{"1 100"}, want: []int32{}},
		{
			name:    "push and pop 1",
			queries: []string{"1 1", "3"},
			want:    []int32{1},
		},
		{
			name:    "push and get 100",
			queries: []string{"1 100", "3"},
			want:    []int32{100},
		},
		{
			name:    "push two ints get one",
			queries: []string{"1 100", "1 2", "3"},
			want:    []int32{100},
		},
		{
			name:    "push two ints, pop one, get one",
			queries: []string{"1 100", "1 2", "2", "3"},
			want:    []int32{100},
		},
		{
			name:    "push two ints, get one",
			queries: []string{"1 100", "1 2", "3"},
			want:    []int32{100},
		},
		{
			name:    "push two, pop max, get one",
			queries: []string{"1 100", "1 102", "2", "3"},
			want:    []int32{100},
		},
		{
			name:    "push three ascending, remove two, get one",
			queries: []string{"1 1", "1 2", "1 3", "2", "2", "3"},
			want:    []int32{1},
		},
		{
			name:    "push three descending, remove two, get one",
			queries: []string{"1 3", "1 2", "1 1", "2", "2", "3"},
			want:    []int32{3},
		},
		{
			name:    "push two, get one, remove one, get one",
			queries: []string{"1 2", "1 3", "3", "2", "3"},
			want:    []int32{3, 2},
		},
		{
			name: "given hackerrank test case",
			queries: []string{
				"1 97",
				"2",
				"1 20",
				"2",
				"1 26",
				"1 20",
				"2",
				"3",
				"1 91",
				"3",
			},
			want: []int32{26, 91},
		},
		{
			name:    "hackerrank test case 5",
			queries: getLines("./testcases/input05.txt"),
			want:    getOutput("./testcases/output05.txt"),
		},
		{
			name:    "hackerrank test case 17",
			queries: getLines("./testcases/input17.txt"),
			want:    getOutput("./testcases/output17.txt"),
		},
	}
	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := getMax(c.queries)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("given %s: got %v want %v", c.name, got, c.want)
			}
		})
	}
}

func getLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var txt []string
	for sc.Scan() {
		txt = append(txt, sc.Text())
	}
	return txt
}

func getOutput(filename string) []int32 {
	var output []int32
	txt := getLines(filename)
	for _, line := range txt {
		n, _ := strconv.ParseInt(line, 10, 32)
		output = append(output, int32(n))
	}
	return output
}
