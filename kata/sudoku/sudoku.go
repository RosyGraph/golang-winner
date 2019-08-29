/*
N-size sudoku is a game with a square table of N2 width and height divided into
N2 smaller squares of N width and height. In a solved state, each of this
smaller squares, as well as each row and column of a full square, contains all
numbers from 1 to N2 without repetition.

Given a number N on the first line and a full sudoku table on the next N2 lines.
Every line contains N2 integers.

Your task is to determine whether this sudoku is solved or not. Output "YES" if
this sudoku table is solved, otherwise "NO".

N can be from 1 to 10.

Sample Input 1:
3
5 8 9 6 7 4 2 1 3
7 4 3 1 8 2 9 5 6
1 2 6 9 5 3 8 7 4
9 3 5 4 2 1 7 6 8
4 1 2 8 6 7 3 9 5
6 7 8 3 9 5 1 4 2
8 6 4 2 1 9 5 3 7
3 9 7 5 4 8 6 2 1
2 5 1 7 3 6 4 8 9
Sample Output 1:
YES

Sample Input 2:
2
1 1 2 2
1 1 2 2
3 3 4 4
3 3 4 4
Sample Output 2:
NO

Sample Input 3:
1
1
Sample Output 3:
YES

Sample Input 4:
2
1 2 3 4
1 2 3 4
1 2 3 4
1 2 3 4

Sample Output 4:
NO

Sample Input 5:
2
1 2 3 4
2 3 4 1
3 4 1 2
4 1 2 3

Sample Output 5:
NO
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func IsSolved(s string) bool {
	rows := strings.Split("\n")
	for _, row := range rows {
		for _, v := range row {
			fmt.Print(v)
		}
	}
	return true
}
