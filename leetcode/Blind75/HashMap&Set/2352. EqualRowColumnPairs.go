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

	for _, col := range grid {
		var colStr string
		for _, r := range col {
			colStr += fmt.Sprintf("%s,", r)
		}

		note[colStr]++
	}

	for row := 0; row < len(grid); row++ {
		var rowStr string
		for col := 0; col < len(grid); col++ {
			rowStr += fmt.Sprintf("%s,", grid[col][row])
		}

		count += note[rowStr]
	}

	return count
}

func equalPairsExercise(grid [][]int) int {
	return 0
}
