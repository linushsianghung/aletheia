package Question_3

// https://leetcode.com/problems/spiral-matrix-ii/description/
/* Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order. */
func generateMatrix(n int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	count := 1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			grid[top][i] = count
			count++
		}
		top++

		for i := top; i <= bottom; i++ {
			grid[i][right] = count
			count++
		}
		right--

		if !(left <= right && top <= bottom) {
			break
		}

		for i := right; i >= left; i-- {
			grid[bottom][i] = count
			count++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			grid[i][left] = count
			count++
		}
		left++
	}

	return grid
}
