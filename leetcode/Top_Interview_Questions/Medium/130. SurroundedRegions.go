package Medium

import "fmt"

// https://leetcode.com/problems/surrounded-regions
/*
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

Connect: A cell is connected to adjacent cells horizontally or vertically.
Region: To form a region connect every 'O' cell.
Surround: The region is surrounded with 'X' cells if you can connect the region with 'X' cells and none of the region cells are on the edge of the board.
A surrounded region is captured by replacing all 'O's with 'X's in the input matrix board.
*/
func solve(board [][]byte) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i > 0 && i < len(board)-1 && j > 0 && j < len(board[0])-1 {
				continue
			}
			exploreBorder(board, i, j, '*', make(map[string]bool))

		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i > 0 && i < len(board)-1 && j > 0 && j < len(board[0])-1 {
				continue
			}
			exploreBorder(board, i, j, 'O', make(map[string]bool))
		}
	}

}

func exploreBorder(board [][]byte, r, c int, marker byte, visited map[string]bool) {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
		return
	}
	if board[r][c] == 'X' {
		return
	}

	position := fmt.Sprint(r, "-", c)
	if visited[position] {
		return
	}
	visited[position] = true
	board[r][c] = marker

	exploreBorder(board, r-1, c, marker, visited)
	exploreBorder(board, r+1, c, marker, visited)
	exploreBorder(board, r, c-1, marker, visited)
	exploreBorder(board, r, c+1, marker, visited)

	return
}
