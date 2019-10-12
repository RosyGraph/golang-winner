/* hackerrank.com algorithm kata */
/* https://www.hackerrank.com/challenges/append-and-delete/problem */

// Problem
////////////////////////////////////////////////////////////////////////////////
/*
 * You have a string of lowercase English alphabetic letters. You can perform two
 * types of operations on the string:
 *     Append a lowercase English alphabetic letter to the end of the string.
 *     Delete the last character in the string. Performing this operation on an
 *     empty string results in an empty string.
 * Given an integer, k, and two strings, s and t, determine whether or not you can
 * convert s to t by performing exactly k of the above operations on s. If it's
 * possible, print Yes. Otherwise, print No.
 *
 * For example, strings s = [a, b, c] and t = [d, e, f]. Our number of moves, k =
 * 6.  To convert s to t, we first delete all of the characters in 3 moves. Next
 * we add each of the characters of t in order. On the 6th move, you will have the
 * matching string. If there had been more moves available, they could have been
 * eliminated by performing multiple deletions on an empty string. If there were
 * fewer than moves, we would not have succeeded in creating the new string.
 */

// Function Description
////////////////////////////////////////////////////////////////////////////////
/*
 * Complete the appendAndDelete function. It should return a string, either Yes or
 * No.
 *
 * appendAndDelete has the following parameters:
 *     s: the initial string
 *     t: the desired string
 *     k: an integer that represents the number of operations
 */

// Input Format
////////////////////////////////////////////////////////////////////////////////
/*
 * The first line contains a string s, the initial string.
 * The second line contains t, the desired final string.
 * The third line contains an integer k, the number of operations.
 */

// Constraints
////////////////////////////////////////////////////////////////////////////////
/*
 * 1 <= |s| <= 100
 * 1 <= |t| <= 100
 * 1 <= |k| <= 100
 */

// Output Format
////////////////////////////////////////////////////////////////////////////////
/* Print Yes if you can obtain the string. Print No otherwise. */
package appdel

import "math"

func appendAndDelete(s, t string, k int32) string {
	var (
		c    int32
		slen int32 = int32(len(s))
		tlen int32 = int32(len(t))
	)

	m := math.Min(float64(slen), float64(tlen))
	for i := 0; i < int(m); i++ {
		if s[i] == t[i] {
			c++
		} else {
			break
		}
	}

	if slen+tlen-2*c > k {
		return "No"
	} else if (slen+tlen-2*c)%2 == k%2 {
		return "Yes"
	} else if slen+tlen-k < 0 {
		return "Yes"
	} else {
		return "No"
	}
}
