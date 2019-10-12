/* hackerrank.com algorithm kata */
/* https://www.hackerrank.com/challenges/append-and-delete/problem */
////////////////////////////////////////////////////////////////////////////////
// Sample 0
////////////////////////////////////////////////////////////////////////////////
// Input
/*
 * hackerhappy
 * hackerrank
 * 9
 */
// Output
/*
 * Yes
 */
// Explaination
/*
 * We perform 5 delete operations to reduce string s to hacker. Next, we perform 4
 * append operations to get hackerrank. Because we were able to convert s to t
 * with k operations, we print Yes.
 */
////////////////////////////////////////////////////////////////////////////////
// Sample 1
////////////////////////////////////////////////////////////////////////////////
// Input
/*
 * aba
 * aba
 * 7
 */
// Output
/*
 * Yes
 */
// Explaination
/*
 * We perform 4 delete operations to reduce string s to the empty string. Next, we
 * perform 3 append operations.
 */
////////////////////////////////////////////////////////////////////////////////
// Sample 2
////////////////////////////////////////////////////////////////////////////////
// Input
/*
 * ashley
 * ash
 * 2
 */
// Output
/*
 * No
 */
// Explaination
/*
 * A minimum of 3 steps are need to convert ashley to ash.
 */

package appdel

import "testing"

func TestAppendAndDelete(t *testing.T) {
	tc := []struct {
		name string
		s    string
		t    string
		k    int32
		want string
	}{
		{name: "sample 0", s: "hackerrank", t: "hackerhappy", k: 9, want: "Yes"},
		{name: "sample 1", s: "aba", t: "aba", k: 7, want: "Yes"},
		{name: "sample 2", s: "ashley", t: "ash", k: 2, want: "No"},
		{name: "sample 3", s: "asb", t: "ash", k: 1, want: "No"},
		{name: "sample 4", s: "qwerasdf", t: "qwerbsdf", k: 6, want: "No"},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := appendAndDelete(c.s, c.t, c.k)
			if got != c.want {
				t.Errorf("%s: got '%s' want '%s'", c.name, got, c.want)
			}
		})
	}
}
