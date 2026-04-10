package Easy

// https://leetcode.com/problems/largest-local-values-in-a-matrix/description/
/*
You are given an n x n integer matrix grid.

Generate an integer matrix maxLocal of size (n - 2) x (n - 2) such that:

maxLocal[i][j] is equal to the largest value of the 3 x 3 matrix in grid centered around row i + 1 and column j + 1.
In other words, we want to find the largest value in every contiguous 3 x 3 matrix in grid.

Return the generated matrix.

Analysis:
Time Complexity: O(n^2)

Because the 3rd and 4th for loop have constant n ("3") so time complexity becomes O(1)
*/
func largestLocal(grid [][]int) [][]int {
	width := len(grid) - 2
	maxLocal := make([][]int, width)
	for i := range maxLocal {
		maxLocal[i] = make([]int, width)
	}

	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			maxNum := 0

			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					maxNum = max(maxNum, grid[k][l])
				}
			}
			maxLocal[i][j] = maxNum
		}
	}

	return maxLocal
}
