package Medium

import "fmt"

// https://leetcode.com/problems/s
//urrounded-regions
// Reference: https://leetcode.com/problems/surrounded-regions/solutions/691675/c-beginner-friendly-boundary-dfs-inplace/
/*
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

Connect: A cell is connected to adjacent cells horizontally or vertically.
Region: To form a region connect every 'O' cell.
Surround: The region is surrounded with 'X' cells if you can connect the region with 'X' cells and none of the region cells are on the edge of the board.
A surrounded region is captured by replacing all 'O's with 'X's in the input matrix board.

Analysis:
Define when an 'O' cannot be flipped:
- If it has at least one 'O' in its adjacent
- The chain of adjacent 'O's is connected to some 'O' which lies on boundary of board

Algorithm:
Using DFS to traverse the boundaries, when 'O' is found, fire the dfs() to flip all the connected 'O' into ' * '.
After that, go through the board and flip the other 'O' to 'X', since these 'O' cannot escape the board. Then flip all the '*' to 'O', indicating these 'O' can reach out of the board.
*/
func solve(board [][]byte) {
	// Step1: Move along the boundary of board, and find O's, Every time we find an O, perform DFS from its position to convert all 'O' to '*'
	note1 := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i > 0 && i < len(board)-1 && j > 0 && j < len(board[0])-1 {
				continue
			}
			// It's required to create its own note
			exploreBorderDFS(board, i, j, '*', note1)
		}
	}

	// Step2. Now board contains three elements,#,O and X. 'O' are left over elements which are not connected to any boundary O, so flip them to 'X'
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	// Step3. '#' are elements which cannot be flipped to 'X', so flip them back to 'O'
	note2 := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if i > 0 && i < len(board)-1 && j > 0 && j < len(board[0])-1 {
				continue
			}
			// It's required to create its own note
			exploreBorderDFS(board, i, j, 'O', note2)
		}
	}

}

func solveExercise(board [][]byte) {
}

func exploreBorderDFS(board [][]byte, r, c int, marker byte, visited map[string]bool) {
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
	exploreBorderDFS(board, r-1, c, marker, visited)
	exploreBorderDFS(board, r+1, c, marker, visited)
	exploreBorderDFS(board, r, c-1, marker, visited)
	exploreBorderDFS(board, r, c+1, marker, visited)
}

func exploreBorderDFSExercise(board [][]byte, r, c int, marker byte, visited map[string]bool) {
}
