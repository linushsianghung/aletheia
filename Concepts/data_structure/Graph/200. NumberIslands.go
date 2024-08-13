package Graph

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-islands/
/*
Given an m x n 2D binary grid 'grid' which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/
func numIslands(grid [][]byte) int {
	count, visited := 0, make(map[string]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if exploreIsland(grid, i, j, visited) {
				count++
			}
		}
	}

	return count
}

func exploreIsland(grid [][]byte, r, c int, visited map[string]bool) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}

	if grid[r][c] == byte('0') {
		return false
	}

	position := fmt.Sprint(r, '-', c)
	if visited[position] {
		return false
	}
	visited[position] = true

	exploreIsland(grid, r+1, c, visited)
	exploreIsland(grid, r-1, c, visited)
	exploreIsland(grid, r, c+1, visited)
	exploreIsland(grid, r, c-1, visited)

	return true
}
