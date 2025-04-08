package pattern

import "fmt"

/*
Write a function, minimumIsland, that takes in a grid containing Ws and Ls. W represents water and L represents land.
The function should return the size of the smallest island. An island is a vertically or horizontally connected region of land.

You may assume that the grid contains at least one island.
*/
func minimumIsland(grid [][]string) int {
	minSize, visited := 0, make(map[string]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			size := exploreMinSize(grid, i, j, visited)
			if minSize > size {
				minSize = size
			}
		}
	}

	return minSize
}

func exploreMinSize(grid [][]string, r, c int, visited map[string]bool) int {
	if r < 0 || r > len(grid) || c < 0 || c > len(grid[0]) {
		return 0
	}
	if grid[r][c] == "W" {
		return 0
	}

	position := fmt.Sprint(r, '-', c)
	if visited[position] {
		return 0
	}
	visited[position] = true

	size := 1
	size += exploreMinSize(grid, r-1, c, visited)
	size += exploreMinSize(grid, r+1, c, visited)
	size += exploreMinSize(grid, r, c-1, visited)
	size += exploreMinSize(grid, r, c+1, visited)
	return size
}
