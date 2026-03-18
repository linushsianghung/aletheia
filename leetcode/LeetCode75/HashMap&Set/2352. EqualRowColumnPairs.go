package HashMap_Set

import "fmt"

// https://leetcode.com/problems/equal-row-and-column-pairs/description/?envId=leetcode-75
/*
Given an n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and a column are considered equal if they contain the same elements in the same order.
*/

func equalPairs(grid [][]int) int {
	// The core idea is to represent each row as a key and count its occurrences.
	// Then, we can construct each column, represent it in the same way, and see how many rows match it.

	// In Go, slices ([]int) cannot be used as map keys because they are not comparable.
	// A common and effective workaround is to convert the slice to a string. fmt.Sprint(slice) creates a string like "[1 2 3]", which is perfect for a key.
	note := make(map[string]int)

	// Step 1: Iterate through each row, convert it to a string, and count its frequency.
	for _, row := range grid {
		// fmt.Sprint is a convenient way to get a consistent string representation of a slice.
		rowKey := fmt.Sprint(row)
		note[rowKey]++
	}

	count := 0
	// Step 2: Iterate through each column.
	for j := range len(grid) {
		// Construct the current column as a slice.
		col := make([]int, len(grid))
		for i := range len(grid) {
			col[i] = grid[i][j]
		}

		// Convert the column to the same string format used for rows.
		colKey := fmt.Sprint(col)

		// If this column's string representation exists in our row map, it means we have found matches.
		// The value at note[colKey] tells us exactly how many rows are identical to this column.
		count += note[colKey]
	}

	return count
}

func equalPairsExercise(grid [][]int) int {
	return 0
}

func equalPairsNaive(grid [][]int) int {
	var count int
	note := make(map[string]int)

	for _, row := range grid {
		var rowStr string
		for _, r := range row {
			rowStr += fmt.Sprintf("%s,", r)
		}

		note[rowStr]++
	}

	for i := range len(grid) {
		var colStr string
		for _, col := range grid {
			colStr += fmt.Sprintf("%s,", col[i])
		}

		count += note[colStr]
	}

	return count
}
