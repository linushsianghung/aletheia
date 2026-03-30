package pattern

import "github.com/linushung/aletheia/Concepts/Data_Structure/Graph"

// IslandCount
/*
Write a function, islandCount, that takes in a grid containing Ws and Ls. W represents water and L represents land.
The function should return the number of islands on the grid. An island is a vertically or horizontally connected region of land.
*/
func IslandCount(grid [][]string) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	// Use a 2D boolean slice for visited tracking. This is much more efficient (O(1) access) than a map with string keys.
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var exploreIsland func(r, c int) bool
	exploreIsland = func(r, c int) bool {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}
		if grid[r][c] == "W" {
			return false
		}

		if visited[r][c] {
			return false
		}
		visited[r][c] = true

		exploreIsland(r+1, c)
		exploreIsland(r-1, c)
		exploreIsland(r, c+1)
		exploreIsland(r, c-1)

		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if exploreIsland(i, j) {
				count++
			}
		}
	}

	return count
}

func IslandCountExercise(grid [][]string) int {
	return 0
}

func IslandCountFlatten(grid [][]string) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	// Optimization: Use a flattened 1D boolean slice.
	// This mimics a 2D array's contiguous memory layout but allows dynamic sizing which requires only 1 allocation instead of (rows + 1) allocations.
	visited := make([]bool, rows*cols)

	var exploreIsland func(grid [][]string, r, c int, visited []bool, cols int) bool
	exploreIsland = func(grid [][]string, r, c int, visited []bool, cols int) bool {
		// Check bounds using correct dimensions: r vs rows, c vs cols
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return false
		}
		if grid[r][c] == "W" {
			return false
		}

		// Map 2D coordinates to 1D index
		flatIndex := r*cols + c
		if visited[flatIndex] {
			return false
		}
		visited[flatIndex] = true

		exploreIsland(grid, r+1, c, visited, cols)
		exploreIsland(grid, r-1, c, visited, cols)
		exploreIsland(grid, r, c+1, visited, cols)
		exploreIsland(grid, r, c-1, visited, cols)

		return true
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if exploreIsland(grid, r, c, visited, cols) {
				count++
			}
		}
	}

	return count
}

// Related Problem: 200. Number of Islands: https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	return Graph.NumIslands(grid)
}
