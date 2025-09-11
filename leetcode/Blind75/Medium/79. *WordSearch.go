package Medium

import "fmt"

// https://leetcode.com/problems/word-search/
// Reference: https://leetcode.com/problems/word-search/solutions/4965052/96-45-easy-solution-with-explanation/
/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.

***Failed in specific case: Time Limit Exceeded ***
*/
func exist(board [][]byte, word string) bool {
	used := make(map[string]bool)

	var searchFunc func(r, c, index int) bool
	searchFunc = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
			return false
		}

		// It's required to put this check after the boarder check, otherwise it will be failed in other Test Case (with the same reason)
		if used[fmt.Sprint(r, "-", c)] {
			return false
		}

		if board[r][c] != word[index] {
			return false
		}

		used[fmt.Sprint(r, "-", c)] = true
		result := searchFunc(r, c+1, index+1) ||
			searchFunc(r, c-1, index+1) ||
			searchFunc(r+1, c, index+1) ||
			searchFunc(r-1, c, index+1)
		used[fmt.Sprint(r, "-", c)] = false

		return result
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if searchFunc(i, j, 0) {
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

	// It's required to put this check after the boarder check, otherwise it will be failed in other Test Case (with the same reason)
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
