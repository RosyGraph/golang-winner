package main

import (
	"fmt"
	"math"
)

func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
	pt := []int32{0, 0}
	d := math.Sqrt(float64(qx-px)*float64(qx-px) + float64(qy-py)*float64(qy-py))
	if px < qx {
		pt[0] = qx + int32(d)
	} else {
		pt[0] = qx - int32(d)
	}

	if py < qy {
		pt[1] = qy + int32(d)
	} else {
		pt[1] = qy - int32(d)
	}

	return pt
}

func main() {
	fmt.Println("vim-go")
}
