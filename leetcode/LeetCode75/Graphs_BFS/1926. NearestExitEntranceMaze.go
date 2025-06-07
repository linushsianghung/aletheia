package Graphs_BFS

import "fmt"

// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/description/?envId=leetcode-75
/*
You are given an m row n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.
In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze. Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.

Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.
*/
func nearestExit(maze [][]byte, entrance []int) int {
	directions := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	note := make(map[string]bool)
	rows, columns, steps := len(maze), len(maze[0]), 0

	queue := [][]int{entrance}
	note[fmt.Sprintf("%d-%d", entrance[0], entrance[1])] = true

	for len(queue) > 0 {
		steps++
		size := len(queue)
		for range size {
			current := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				row := current[0] + direction[0]
				column := current[1] + direction[1]

				// Below order of validation is really really matter!!!
				// If current position goes beyond the maze, just skip to next direction
				if row < 0 || row == rows || column < 0 || column == columns {
					continue
				}
				// It could be out of range if validating it first
				if maze[row][column] == byte('+') {
					continue
				}
				// It has to be validated before next one, otherwise it might finish quite quick if the entrance is already on the boarder
				position := fmt.Sprintf("%d-%d", row, column)
				if ok := note[position]; ok {
					continue
				}
				// If current position reaches the boarder/exit, return steps
				if row == 0 || row == rows-1 || column == 0 || column == columns-1 {
					return steps
				}

				note[position] = true
				queue = append(queue, []int{row, column})
			}
		}
	}

	return -1
}

// Related Problem: 994. Rotting Oranges: https://leetcode.com/problems/rotting-oranges/description/
func orangesRottingProblem(grid [][]int) int {
	return orangesRotting(grid)
}
