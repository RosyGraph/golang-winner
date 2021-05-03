package main

import "testing"

func TestEqualStacks(t *testing.T) {
	tests := []struct {
		name       string
		h1, h2, h3 []int32
		want       int32
	}{
		{
			name: "empty stacks",
			h1:   []int32{},
			h2:   []int32{},
			h3:   []int32{},
			want: 0,
		},
		{
			name: "equal stacks",
			h1:   []int32{1},
			h2:   []int32{1},
			h3:   []int32{1},
			want: 1,
		},
		{
			name: "one stack too high",
			h1:   []int32{1, 1},
			h2:   []int32{1},
			h3:   []int32{1},
			want: 1,
		},
		{
			name: "two stacks too high (1)",
			h1:   []int32{1, 1},
			h2:   []int32{1, 1, 1},
			h3:   []int32{1, 1, 1},
			want: 2,
		},
		{
			name: "two stacks too high (2)",
			h1:   []int32{1, 1, 1, 1},
			h2:   []int32{1, 1, 1, 1},
			h3:   []int32{1, 1, 1},
			want: 3,
		},
		{
			name: "two stacks too high (3)",
			h1:   []int32{1, 1, 1, 1, 1},
			h2:   []int32{1, 1, 1, 1},
			h3:   []int32{1, 1, 1, 1, 1},
			want: 4,
		},
		{
			name: "given testcase",
			h1:   []int32{3, 2, 1, 1, 1},
			h2:   []int32{4, 3, 2},
			h3:   []int32{1, 1, 4, 1},
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := equalStacks(test.h1, test.h2, test.h3)
			if got != test.want {
				t.Errorf("test [%s]:\ngot %d want %d\n", test.name, got, test.want)
			}
		})
	}
}
