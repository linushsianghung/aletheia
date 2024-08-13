package pattern

import "fmt"

// IslandCount
/*
Write a function, islandCount, that takes in a grid containing Ws and Ls. W represents water and L represents land.
The function should return the number of islands on the grid. An island is a vertically or horizontally connected region of land.

Related Topic: 200. Number of Islands: https://leetcode.com/problems/number-of-islands/
*/
func IslandCount(grid [][]string) int {
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

func exploreIsland(grid [][]string, r, c int, visited map[string]bool) bool {
	if r < 0 || r >= len(grid[0]) || c < 0 || c >= len(grid) {
		return false
	}

	if grid[r][c] == "W" {
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
