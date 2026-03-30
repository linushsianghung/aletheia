package pattern

/*
Write a function, minimumIsland, that takes in a grid containing Ws and Ls. W represents water and L represents land.
The function should return the size of the smallest island. An island is a vertically or horizontally connected region of land.

You may assume that the grid contains at least one island.
*/
func minimumIsland(grid [][]string) int {
	rows, cols := len(grid), len(grid[0])
	// Initialize minSize to the largest possible size. This ensures that the size of the first island found will be smaller.
	minSize := rows * cols

	// Optimization: Use a flattened 1D slice for the visited set. This improves performance by using a single memory allocation and improving cache locality.
	visited := make([]bool, rows*cols)

	var exploreMinSize func(r, c int) int
	exploreMinSize = func(r, c int) int {
		// 1. Check bounds: Ensure row and column are within the grid.
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return 0
		}
		// 2. Check for water: We cannot traverse water.
		if grid[r][c] == "W" {
			return 0
		}

		// 3. Check if already visited to prevent infinite loops and recounting.
		flatIndex := r*cols + c
		if visited[flatIndex] {
			return 0
		}
		visited[flatIndex] = true

		// The size of the current island component is 1 (for this cell) plus the size of any connected land cells.
		return 1 + exploreMinSize(r-1, c) + exploreMinSize(r+1, c) + exploreMinSize(r, c-1) + exploreMinSize(r, c+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			size := exploreMinSize(i, j)
			// Only update minSize if we've found a new island (size > 0) and its size is smaller than the current minimum.
			if size > 0 && size < minSize {
				minSize = size
			}
		}
	}

	return minSize
}

func minimumIslandExercise(grid [][]string) int {
	return 0
}
