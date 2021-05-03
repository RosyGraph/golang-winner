package main

import "testing"

func TestBalanced(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single bracket",
			input: "(",
			want:  "NO",
		},
		{
			name:  "one set of matching brackets",
			input: "[]",
			want:  "YES",
		},
		{
			name:  "nested brackets",
			input: "[({{()}})]",
			want:  "YES",
		},
		{
			name:  "mismatched nested brackets",
			input: "[({{(})})]",
			want:  "NO",
		},
		{
			name:  "too few opening brackets",
			input: "[({{})})]",
			want:  "NO",
		},
		{
			name:  "too few closing brackets",
			input: "{{(([[[]]))}}",
			want:  "NO",
		},
		{
			name:  "leading closing bracket",
			input: "}",
			want:  "NO",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isBalanced(test.input)
			if got != test.want {
				t.Errorf("given %s:\n\tgot: %s, want:%s", test.name, got, test.want)
			}
		})
	}
}
