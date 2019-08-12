package kata

import "testing"

func TestLongestConsec(t *testing.T) {
	testCases := []struct {
		name   string
		strarr []string
		want   string
		k      int
	}{
		{
			name:   "lower case looking at 2 words",
			strarr: []string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"},
			want:   "abigailtheta",
			k:      2,
		},
		{
			name:   "zero length string array",
			strarr: []string{},
			want:   "",
			k:      0,
		},
		{
			name:   "k > length of string array",
			strarr: []string{"hello", "world"},
			want:   "",
			k:      4,
		},
		{
			name:   "basic case",
			strarr: []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
			want:   "oocccffuucccjjjkkkjyyyeehh",
			k:      1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsec(tt.strarr, tt.k)

			if tt.want != got {
				t.Errorf("got %s want %s", got, tt.want)
			}
		})
	}
}
