package Question_4

// https://leetcode.com/problems/number-of-enclaves/description/
/*
You are given an m x n binary matrix grid, where 0 represents a sea cell and 1 represents a land cell.

A move consists of walking from one land cell to another adjacent (4-directionally) land cell or walking off the boundary of the grid.

Return the number of land cells in grid for which we cannot walk off the boundary of the grid in any number of moves.
*/
func numEnclaves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	// In this scenario, it's unnecessary to use note to record the traversal because we flip the island first and check if it's water in every dfs
	//note := make([][]bool, rows)
	//for i := range note {
	//	note[i] = make([]bool, cols)
	//}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if grid[r][c] == 0 {
			return
		}

		grid[r][c] = 0

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > 0 && i < rows-1 && j > 0 && j < cols-1 {
				continue
			}
			if grid[i][j] == 0 {
				continue
			}

			dfs(i, j)
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	return count
}
