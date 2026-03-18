package Medium

// https://leetcode.com/problems/set-matrix-zeroes/
/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
*/
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	flipped := make([][]bool, rows)
	for i := range flipped {
		flipped[i] = make([]bool, cols)
	}

	var setZeroFunc func(r, c int)
	setZeroFunc = func(r, c int) {
		for i := range matrix[r] {
			if matrix[r][i] == 0 {
				continue
			}
			matrix[r][i] = 0
			flipped[r][i] = true
		}

		for i, row := range matrix {
			if row[c] == 0 {
				continue
			}

			row[c] = 0
			flipped[i][c] = true
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 && !flipped[i][j] {
				setZeroFunc(i, j)
			}
		}
	}
}

func setZeroesExercise(matrix [][]int) {

}

func setZeroFunc(matrix [][]int, r, c int, flipped [][]bool) {
	for i := range matrix[r] {
		if matrix[r][i] == 0 {
			continue
		}
		matrix[r][i] = 0
		flipped[r][i] = true
	}

	for i, row := range matrix {
		if row[c] == 0 {
			continue
		}

		row[c] = 0
		flipped[i][c] = true
	}
}
