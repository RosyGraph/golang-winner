package main

import (
	"reflect"
	"testing"
)

func TestReadAsStringArray(t *testing.T) {
	tc := []struct {
		name     string
		filename string
		want     []string
	}{
		{
			name:     "hello world",
			filename: "dat/test1.txt",
			want:     []string{"hello", "world"},
		},
		{
			name:     "HeLlO WoRlD",
			filename: "dat/test2.txt",
			want:     []string{"hello", "world"},
		},
		{
			name:     "HELLO WORLD!",
			filename: "dat/test3.txt",
			want:     []string{"hello", "world"},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := ReadAsStringArray(c.filename)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("filename %s got '%v' want '%v'",
					c.filename, got, c.want)
			}
		})
	}
}

func TestClean(t *testing.T) {
	tc := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "one capital letter",
			s:    "Hello",
			want: "hello",
		},
		{
			name: "one punctuation symbol",
			s:    "aardvark!",
			want: "aardvark",
		},
		{
			name: "one punctuation symbol and one capital letter",
			s:    "aNcill4ary",
			want: "ancillary",
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := Clean(c.s)

			if got != c.want {
				t.Errorf("got %s want %s", got, c.want)
			}
		})

	}
}

func TestContains(t *testing.T) {
	tc := []struct {
		name string
		arr  []string
		s    string
		want bool
	}{
		{
			name: "array contains one instance",
			arr:  []string{"hello", "world", "z", "zz"},
			s:    "world",
			want: true,
		},
		{
			name: "array contains no instance",
			arr:  []string{"a", "b", "c", "c"},
			s:    "aardvark",
			want: false,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := Contains(c.arr, c.s)

			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}
		})
	}
}

func TestCountOccurences(t *testing.T) {
	tc := []struct {
		name         string
		keys, values []string
		want         int
	}{
		{
			name:   "4 occurences",
			keys:   []string{"d", "e", "f", "", "f", "g"},
			values: []string{"a", "b", "c", "d", "e", "f"},
			want:   4,
		},
		{
			name:   "0 occurences",
			keys:   []string{"", "", "", "", "", ""},
			values: []string{"a", "b", "c", "d", "e", "f"},
			want:   0,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := CountOccurences(c.keys, c.values)

			if got != c.want {
				t.Errorf("got %d want %d", got, c.want)
			}
		})
	}
}

func TestCalculateSentiment(t *testing.T) {
	tc := []struct {
		name string
		arr  []string
		want float64
	}{
		{
			name: "all negative sentiment",
			arr:  []string{"bad", "ugly", "sad"},
			want: -100.0,
		},
		{
			name: "all positive sentiment",
			arr:  []string{"happy", "good", "great"},
			want: 100.0,
		},
		{
			name: "half positive half negative sentiment",
			arr:  []string{"happy", "good", "bad", "terrible"},
			want: 0.0,
		},
		{
			name: "half positive half negative sentiment with stop words",
			arr:  []string{"happy", "good", "us", "bad", "terrible", "we"},
			want: 0.0,
		},
	}

	for _, c := range tc {
		t.Run("all negative sentiment", func(t *testing.T) {
			got := CalculateSentiment(c.arr)

			if got != c.want {
				t.Errorf("got %f want %f", got, c.want)
			}
		})
	}
}
