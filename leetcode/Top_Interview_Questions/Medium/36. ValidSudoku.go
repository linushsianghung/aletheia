package Medium

import "fmt"

// https://leetcode.com/problems/valid-sudoku/
/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3.Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.
*/
func isValidSudoku(board [][]byte) bool {
	note := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			numInRow := fmt.Sprintf("Row %d has %d", i, board[i][j])
			if _, ok := note[numInRow]; ok {
				return false
			}
			numInColumn := fmt.Sprintf("Column %d has %d", j, board[i][j])
			if _, ok := note[numInColumn]; ok {
				return false
			}
			subBox := fmt.Sprintf("%d-%d", i/3, j/3)
			numInSubBox := fmt.Sprintf("Subbod %s has %d", subBox, board[i][j])
			if _, ok := note[numInSubBox]; ok {
				return false
			}

			note[numInRow] = true
			note[numInColumn] = true
			note[numInSubBox] = true
		}
	}

	return true
}
