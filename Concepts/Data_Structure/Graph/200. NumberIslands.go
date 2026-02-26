package Graph

// NumIslands https://leetcode.com/problems/number-of-islands/
/*
Given an m x n 2D binary grid 'grid' which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/
func NumIslands(grid [][]byte) int {
	count, visited := 0, make([][]bool, len(grid))

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if exploreIsland(grid, row, col, visited) {
				count++
			}
		}
	}

	return count
}

func numIslandsExercise(grid [][]byte) int {
	return 0
}

func exploreIsland(grid [][]byte, r, c int, visited [][]bool) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}
	if grid[r][c] == byte('0') {
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
