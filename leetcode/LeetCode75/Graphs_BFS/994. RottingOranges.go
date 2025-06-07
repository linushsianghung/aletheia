package Graphs_BFS

import "fmt"

// https://leetcode.com/problems/rotting-oranges/description/?envId=leetcode-75
/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
*/
func orangesRotting(grid [][]int) int {
	directions := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	rows, columns := len(grid), len(grid[0])
	note := make(map[string]bool)

	// Because of a special test case [[0]] which expects 0, using a variable to indicate fresh oranges in order to distinguish another test case [[1]] which expects -1
	fresh := 0
	// Find out all the existing rotting oranges
	queue := make([][]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	// To handle the special test case [[0]] which expects 0
	if len(queue) == 0 && fresh == 0 {
		return 0
	}

	minutes := -1
	for len(queue) > 0 {
		minutes++
		size := len(queue)

		for range size {
			current := queue[0]
			queue = queue[1:]
			position := fmt.Sprintf("%d-%d", current[0], current[1])
			note[position] = true

			for _, direction := range directions {
				row := current[0] + direction[0]
				column := current[1] + direction[1]

				// Below order of validation is really really matter!!!
				// If current position goes beyond the maze, just skip to next direction
				if row < 0 || row == rows || column < 0 || column == columns {
					continue
				}
				// It could be out of range if validating it first
				if grid[row][column] == 0 || grid[row][column] == 2 {
					continue
				}

				position = fmt.Sprintf("%d-%d", row, column)
				if ok := note[position]; ok {
					continue
				}

				fresh--
				grid[row][column] = 2
				queue = append(queue, []int{row, column})
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return minutes
}

// Related Problem: 1926. Nearest Exit from Entrance in Maze: https://leetcode.com/problems/rotting-oranges/description
func nearestExitProblem(maze [][]byte, entrance []int) int {
	return nearestExit(maze, entrance)
}
