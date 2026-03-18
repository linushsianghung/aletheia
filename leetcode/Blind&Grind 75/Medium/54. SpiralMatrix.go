package Medium

// https://leetcode.com/problems/spiral-matrix/
// Reference: https://neetcode.io/solutions/spiral-matrix
/* Given an m x n matrix, return all elements of the matrix in spiral order. */
func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix[0])

	for left < right {
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right--

		if left == right || top == bottom {
			break
		}

		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom-1][i])
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++

		if left == right || top == bottom {
			break
		}
	}

	return result
}
