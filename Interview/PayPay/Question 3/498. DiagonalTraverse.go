package Question_3

// https://leetcode.com/problems/diagonal-traverse/description/
/* Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order. */
func findDiagonalOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, rows*cols)

	row, col := 0, 0
	for i := 0; i < rows*cols; i++ {
		result[i] = matrix[row][col]

		// If (row + col) is even → move up-right (row--, col++)
		if (row+col)%2 == 0 {
			// If meet first row & last column at the same time, it should change the direction to move down (row++)
			if col == cols-1 {
				row++
			} else if row == 0 {
				col++
			} else {
				row--
				col++
			}
			// If (row + col) is odd → move down-left (row++, col--)
		} else {
			// If meet first col & last row at the same time, it should change the direction to move right (col++)
			if row == rows-1 {
				col++
			} else if col == 0 {
				row++
			} else {
				row++
				col--
			}
		}
	}

	return result
}
