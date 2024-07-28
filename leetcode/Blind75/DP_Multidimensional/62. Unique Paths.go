package DP_Multidimensional

import (
	"github.com/linushung/aletheia/leetcode/Top_Interview_Questions/Medium"
)

// UniquePaths https://leetcode.com/problems/unique-paths/description
/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 10^9.
*/
func UniquePaths(m int, n int) int {

	// return Medium.UniquePaths(m, n)
	return Medium.UniquePathsMemorisation(m, n, make(map[string]int))
}
