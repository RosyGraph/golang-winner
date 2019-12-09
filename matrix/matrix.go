package matrix

type Matrix [][]int

func NewMatrix(n, m int) Matrix {
	matrix := make(Matrix, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	return matrix
}
