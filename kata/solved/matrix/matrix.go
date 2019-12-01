/*
Calculates highest 'hourglass' sum in a given 2d array
*/
package main

import "math"

func hourglassSum(a [][]int32) int32 {
	var max int32 = math.MinInt32
	for y := 1; y < len(a)-1; y++ {
		for x := 1; x < len(a[0])-1; x++ {
			if s := sumGlass(a, x, y); s > max {
				max = s
			}
		}
	}
	return max
}

func sumGlass(a [][]int32, x, y int) int32 {
	return a[y-1][x-1] + a[y-1][x] + a[y-1][x+1] +
		a[y][x] + a[y+1][x-1] + a[y+1][x] + a[y+1][x+1]
}
