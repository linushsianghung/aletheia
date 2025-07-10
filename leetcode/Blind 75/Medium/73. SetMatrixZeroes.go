package Medium

import "fmt"

// https://leetcode.com/problems/set-matrix-zeroes/
/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
*/
func setZeroes(matrix [][]int) {
	flipped := make(map[string]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// Only deal with 0 and not flipped by purposed
			if matrix[i][j] == 0 && !flipped[fmt.Sprint(i, "-", j)] {
				setZeroFunc(matrix, i, j, "right", flipped)
				setZeroFunc(matrix, i, j, "left", flipped)
				setZeroFunc(matrix, i, j, "up", flipped)
				setZeroFunc(matrix, i, j, "down", flipped)
			}
		}
	}
}

func setZeroesExercise(matrix [][]int) {

}

func setZeroFunc(matrix [][]int, i, j int, direction string, flipped map[string]bool) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return
	}

	// If the element not 0, flip it! And record that it is flipped by purposed
	if matrix[i][j] != 0 {
		matrix[i][j] = 0
		flipped[fmt.Sprint(i, "-", j)] = true
	}

	switch direction {
	case "right":
		setZeroFunc(matrix, i, j+1, "right", flipped)
	case "left":
		setZeroFunc(matrix, i, j-1, "left", flipped)
	case "up":
		setZeroFunc(matrix, i-1, j, "up", flipped)
	case "down":
		setZeroFunc(matrix, i+1, j, "down", flipped)
	}
}

func setZeroFuncExercise(matrix [][]int, i, j int, direction string, flipped map[string]bool) {
}
