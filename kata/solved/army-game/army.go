package army

import "math"

func gameWithCells(n, m int32) int32 {
	return int32(math.Ceil(float64(n)/2.0) * math.Ceil(float64(m)/2.0))
}

func MinDrops(bases []int) int {
	return int(math.Ceil(float64(bases[0])/2.0) * math.Ceil(float64(bases[1])/2.0))
}
