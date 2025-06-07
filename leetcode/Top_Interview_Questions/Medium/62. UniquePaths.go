package Medium

import (
	"fmt"
	"github.com/linushung/aletheia/Concepts/algorithm/DynamicProgramming"
)

// UniquePaths https://leetcode.com/problems/unique-paths/description/
// Ref:
// - https://leetcode.com/problems/unique-paths/solutions/22954/c-dp/
// - https://leetcode.com/problems/unique-paths/solutions/1581998/c-python-5-simple-solutions-w-explanation-optimization-from-brute-force-to-dp-to-math/
/*
There is a robot on an m x n grid. The robot is initially located in the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.
*/
func UniquePaths(m int, n int) int {

	return UniquePathsMemorisation(m, n, make(map[string]int))
	// return uniquePaths2Slice(m, n)
	//return uniquePathsSmartSlice(m, n)
}

func UniquePathsMemorisation(m, n int, memo map[string]int) int {
	key := fmt.Sprint(m, '-', n)
	if value, ok := memo[key]; ok {
		return value
	}
	// Here the "or" operator can do the job for improvement because whenever reaching the edge, there is definitely one way to arrive the end point
	if m == 1 || n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	memo[key] = UniquePathsMemorisation(m-1, n, memo) + UniquePathsMemorisation(m, n-1, memo)
	return memo[key]
}

// Using Backtracking strategy as practice and as expect it will get "Memory Limit Exceeded" error message when the grid become larger
func uniquePathsBacktracking(m, n int) int {
	result := make([][]string, 0)

	var nestedFunc func(processor []string, m, n int)
	nestedFunc = func(processor []string, m, n int) {
		if m == 0 || n == 0 {
			return
		}

		if m == 1 && n == 1 {
			result = append(result, processor)
		}

		if m > 0 {
			nestedFunc(append(processor, "down"), m-1, n)
		}
		if n > 0 {
			nestedFunc(append(processor, "right"), m, n-1)
		}
	}
	nestedFunc(make([]string, 0), m, n)
	return len(result)
}

/*
[1, 1, 1,  1,  1,  1,  1]
=> [1, (1 + 1), (1 + 2), (1 + 3), (1 + 4), (1 + 5), (1 + 6)] = [1, 2, 3, 4, 5, 6, 7]

[1, 2, 3, 4, 5, 6, 7]
=> [1, (2 + 1), (3 + 3), (4 + 6), (5 + 10), (6 + 15), (7 + 21)] = [1, 3, 6, 10, 15, 21, 28]

Further observation that for each new iteration, pre[j] is just the cur[j] before the update.
*/
func uniquePathsSmartSlice(m int, n int) int {
	cur := make([]int, n)
	// Initialise table
	for i := 0; i < n; i++ {
		cur[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = cur[j] + cur[j-1]
		}
	}

	return cur[n-1]
}

/*
[1, 1, 1,  1,  1,  1,  1]
[1, 2, 3,  4,  5,  6,  7]
[1, 3, 6, 10, 15, 21, 28]
*/
func uniquePaths2Slice(m int, n int) int {
	pre, cur := make([]int, n), make([]int, n)
	// Initialise slice
	for i := 0; i < n; i++ {
		pre[i] = 1
	}
	// As long as we initialise the first element as 1, it doesn't matter what number are for the rest of number in slice
	cur[0] = 1

	// Using cur slice for calculating the result and then store the result to pre slice
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = pre[j] + cur[j-1]
		}

		// Store the result so we can calculate for next iteration by using cur slice again
		pre = cur
	}

	return pre[n-1]
}

func uniquePathsBF(m int, n int) int {
	// Standard way to create 2D slice
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// Put 1 to the first row
	for _, col := range grid {
		col[0] = 1
	}

	// Put 1 to the first column
	for i := range grid[0] {
		grid[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}

// Related Problem: 63. Unique Paths II: https://leetcode.com/problems/unique-paths-ii/description/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return DynamicProgramming.UniquePathsWithObstacles(obstacleGrid)
}
