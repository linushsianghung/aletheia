package Medium

// https://leetcode.com/problems/word-search/
// Reference: https://leetcode.com/problems/word-search/solutions/4965052/96-45-easy-solution-with-explanation/
/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.
*/
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	if len(word) > rows*cols {
		return false
	}

	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}
		if board[r][c] != word[index] {
			return false
		}

		// Mark as visited by temporarily changing the board content
		temp := board[r][c]
		board[r][c] = '#'

		// Explore neighbors
		found := dfs(r+1, c, index+1) ||
			dfs(r-1, c, index+1) ||
			dfs(r, c+1, index+1) ||
			dfs(r, c-1, index+1)

		// Backtrack: Restore the original character
		board[r][c] = temp
		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Optimization: Only start DFS if the first character matches
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func existExercise(board [][]byte, word string) bool {
	return false
}

func existNaive(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var existFunc func(r, c, index int, visited [][]bool) bool
	existFunc = func(r, c, index int, visited [][]bool) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}
		if board[r][c] != word[index] {
			return false
		}
		if visited[r][c] {
			return false
		}
		visited[r][c] = true

		isExist := existFunc(r+1, c, index+1, visited) ||
			existFunc(r-1, c, index+1, visited) ||
			existFunc(r, c+1, index+1, visited) ||
			existFunc(r, c-1, index+1, visited)
		if isExist {
			return true
		}

		visited[r][c] = false
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visited := make([][]bool, rows)
			for i := range visited {
				visited[i] = make([]bool, cols)
			}

			if existFunc(i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}
