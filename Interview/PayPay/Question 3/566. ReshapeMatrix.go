package Question_3

// https://leetcode.com/problems/reshape-the-matrix/description/
/*
In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.

You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.

The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.
*/
func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows*cols != r*c {
		return mat
	}

	reshape := make([][]int, r)
	for i := range reshape {
		reshape[i] = make([]int, c)
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			reshape[count/c][count%c] = mat[i][j]
			count++
		}
	}

	return reshape
}

func matrixReshapeNaive(mat [][]int, r int, c int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows*cols != r*c {
		return mat
	}

	flatten := make([]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			flatten = append(flatten, mat[i][j])
		}
	}

	grid := make([][]int, r)
	for i := range grid {
		grid[i] = make([]int, c)
	}
	count := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			grid[i][j] = flatten[count]
			count++
		}
	}

	return grid
}
