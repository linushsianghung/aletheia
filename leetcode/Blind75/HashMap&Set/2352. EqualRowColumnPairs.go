package HashMap_Set

import "fmt"

// https://leetcode.com/problems/equal-row-and-column-pairs/description/?envId=leetcode-75
/*
Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).
*/
func equalPairs(grid [][]int) int {
	var count int
	note := make(map[string]int)

	for _, row := range grid {
		var rowStr string
		for _, c := range row {
			rowStr += fmt.Sprintf("%s,", c)
		}

		note[rowStr]++
	}

	for col := 0; col < len(grid); col++ {
		var colStr string
		for row := 0; row < len(grid); row++ {
			colStr += fmt.Sprintf("%s,", grid[row][col])
		}

		count += note[colStr]
	}

	return count
}

func equalPairsExercise(grid [][]int) int {
	return 0
}
