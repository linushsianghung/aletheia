package Easy

// https://leetcode.com/problems/toeplitz-matrix/description/
/*
Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return false.

A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same elements.
*/
func isToeplitzMatrix(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}

func isToeplitzMatrixExercise(matrix [][]int) bool {
	return false
}
