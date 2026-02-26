package Graph

// https://leetcode.com/problems/max-area-of-island/
/*
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.
*/
func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return 0
		}
		if grid[r][c] == 0 {
			return 0
		}
		if visited[r][c] {
			return 0
		}
		visited[r][c] = true

		return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
	}

	maxCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				continue
			}

			count := dfs(i, j)
			maxCount = max(maxCount, count)
		}
	}

	return maxCount
}
