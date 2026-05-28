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

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 && !flipped[i][j] {
				for row := range rows {
					if matrix[row][j] != 0 {
						matrix[row][j] = 0
						flipped[row][j] = true
					}
				}

				for col := range cols {
					if matrix[i][col] != 0 {
						matrix[i][col] = 0
						flipped[i][col] = true
					}
				}
			}
		}
	}
}

func setZeroesExercise(matrix [][]int) {

}
