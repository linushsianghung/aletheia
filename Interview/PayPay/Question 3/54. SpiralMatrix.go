package Question_3

// SpiralOrder https://leetcode.com/problems/spiral-matrix/
// Reference: https://neetcode.io/solutions/spiral-matrix
/*
Given an m x n matrix, return all elements of the matrix in spiral order.

Analysis:


*/
func SpiralOrder(matrix [][]int) []int {
	return spiralOrderClosedIntervals(matrix)
}

func spiralOrderClosedIntervals(matrix [][]int) []int {
	result := make([]int, 0)
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if !(left <= right && top <= bottom) {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
}

func spiralOrderOpenHalf(matrix [][]int) []int {
	result := make([]int, 0)
	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix[0])

	for top < bottom && left < right {
		// 1. Traverse Top Row (left → right)
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// 2. Traverse Right Column (top → bottom)
		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right--

		// It is needed because after shrinking boundaries and end up with a single row: Prevent duplicate traversal
		if !(top < bottom && left < right) {
			break
		}

		// 3. Traverse Bottom Row (right → left)
		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom-1][i])
		}
		bottom--

		// 4. Traverse Left Column (bottom → top)
		for i := bottom - 1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
}
