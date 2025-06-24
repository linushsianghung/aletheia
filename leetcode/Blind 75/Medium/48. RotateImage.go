package Medium

// https://leetcode.com/problems/rotate-image/
// Reference: NeetCode: https://www.youtube.com/watch?v=fMSJSS7eO1wｐｐ
/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
ｐ
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
*/
func rotate(matrix [][]int) {
	left, right := 0, len(matrix)-1

	for left < right {
		for i := 0; i < right-left; i++ {
			top, bottom := left, right

			// store the Top-Left
			temp := matrix[top][left+i]

			// move Bottom-Left to Top-Left
			matrix[top][left+i] = matrix[bottom-i][left]

			// move Bottom-Right to Bottom-Left
			matrix[bottom-i][left] = matrix[bottom][right-i]

			// move Top-Right to Bottom-Right
			matrix[bottom][right-i] = matrix[top+i][right]

			// move Top-Left to Top-Right
			matrix[top+i][right] = temp
		}

		left++
		right--
	}
}

func rotateExercise(matrix [][]int) {
	left, right := 0, len(matrix)

	for left < right {
		for i := 0; i < right-left; i++ {
			top, bottom := left, right

			temp := matrix[top][left+i]
			matrix[top][left+i] = matrix[bottom-i][left]
			matrix[bottom-i][left] = matrix[bottom][right-i]
			matrix[bottom][right-i] = matrix[top+i][right]
			matrix[top+i][right] = temp
		}

		left++
		right--
	}

}
