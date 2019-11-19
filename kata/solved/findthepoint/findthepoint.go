package main

import (
	"fmt"
	"math"
)

func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
	pt := []int32{0, 0}
	dx := int32(math.Sqrt(float64(px-qx) * float64(px-qx)))
	dy := int32(math.Sqrt(float64(py-qy) * float64(py-qy)))

	if px < qx {
		pt[0] = qx + dx
	} else {
		pt[0] = qx - dx
	}

	if py < qy {
		pt[1] = qy + dy
	} else {
		pt[1] = qy - dy
	}

	return pt
}

func main() {
	fmt.Println("vim-go")
}
