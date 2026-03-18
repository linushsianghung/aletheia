package Binary_Search

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/description/
/*
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.
*/
func countNegatives(grid [][]int) int {
	count, n := 0, len(grid[0])
	for i := 0; i < len(grid); i++ {
		left, right := 0, n-1

		for left <= right {
			mid := left + (right-left)/2
			if grid[i][mid] >= 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		count += n - left
	}

	return count
}

func countNegativesExercise(grid [][]int) int {
	return 0
}
