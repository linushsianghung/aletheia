package Medium

import "fmt"

// https://leetcode.com/problems/word-search/
/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.
*/
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if searchFunc(board, i, j, 0, word, make(map[string]bool)) {
				return true
			}
		}
	}

	return false
}

func searchFunc(board [][]byte, i, j, index int, word string, used map[string]bool) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if used[fmt.Sprint(i, "-", j)] {
		return false
	}

	if board[i][j] != word[index] {
		return false
	}

	used[fmt.Sprint(i, "-", j)] = true
	result := searchFunc(board, i, j+1, index+1, word, used) ||
		searchFunc(board, i, j-1, index+1, word, used) ||
		searchFunc(board, i+1, j, index+1, word, used) ||
		searchFunc(board, i-1, j, index+1, word, used)
	used[fmt.Sprint(i, "-", j)] = false

	return result
}
