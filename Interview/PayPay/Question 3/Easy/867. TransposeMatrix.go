package Easy

// https://leetcode.com/problems/transpose-matrix/description/
// Reference: https://leetcode.com/problems/transpose-matrix/solutions/6097971/video-give-me-5-minutes-2-solutions-how-ihrf8/
/*
Given a 2D integer array matrix, return the transpose of matrix.

The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
*/
func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])

	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := range rows {
		for j := range cols {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

func transposeExercise(matrix [][]int) [][]int {
	return nil
}
