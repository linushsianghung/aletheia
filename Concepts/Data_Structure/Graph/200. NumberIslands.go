package Graph

// NumIslands https://leetcode.com/problems/number-of-islands/
/*
Given an m x n 2D binary grid 'grid' which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/
func NumIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	note := make([][]bool, rows)
	for i := range note {
		note[i] = make([]bool, cols)
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Optimization: Only start to explore if the grid is land
			if grid[i][j] != '0' && exploreIsland(grid, i, j, note) {
				count++
			}
		}
	}

	return count
}

func exploreIsland(grid [][]byte, r, c int, visited [][]bool) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}
	if grid[r][c] == '0' {
		return false
	}

	if visited[r][c] {
		return false
	}
	visited[r][c] = true

	exploreIsland(grid, r+1, c, visited)
	exploreIsland(grid, r-1, c, visited)
	exploreIsland(grid, r, c+1, visited)
	exploreIsland(grid, r, c-1, visited)

	return true
}

func numIslandsExercise(grid [][]byte) int {
	return 0
}

func exploreIslandExercise(grid [][]byte, r, c int, visited [][]bool) bool {
	return false
}
