package Medium

import "github.com/linushung/aletheia/Concepts/algorithm/Backtracking"

// https://leetcode.com/problems/generate-parentheses/description/
func generateParenthesis(n int) []string {
	return Backtracking.GenerateParenthesis(n)
}
