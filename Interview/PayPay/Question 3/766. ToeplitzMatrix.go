package Question_3

// https://leetcode.com/problems/toeplitz-matrix/description/
/*
Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return false.

A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same elements.
*/
func isToeplitzMatrix(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == rows-1 || j == cols-1 {
				continue
			}

			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}
