package DynamicProgramming

// https://leetcode.com/problems/pascals-triangle/
/*
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Analysis:
		1
	  1   1
	1   2   1
  1   3   3   1
1   4   6   4   1

# Using 2D slice to create Pascals Triangle
triangle := [row][element]
# Based on the definition to sum up the each element
triangle[i][j] = preRow[j-1] + preRow[j]
*/
func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	if numRows == 0 {
		return triangle
	}

	return pascalsTriangle(numRows, triangle)
}

func pascalsTriangle(numRows int, triangle [][]int) [][]int {
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}

	triangle[0][0] = 1
	for i := 1; i < numRows; i++ {
		// Set left edge element
		triangle[i][0] = 1

		preRow := triangle[i-1]
		for j := 1; j < i; j++ {
			triangle[i][j] = preRow[j-1] + preRow[j]
		}

		// Set right edge element
		triangle[i][i] = 1
	}

	return triangle
}

func pascalsTriangleExercise(numRows int, triangle [][]int) [][]int {
	return triangle
}
