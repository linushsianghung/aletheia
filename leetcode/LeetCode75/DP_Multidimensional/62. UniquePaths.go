package DP_Multidimensional

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions/Medium"

// https://leetcode.com/problems/unique-paths/description/?envId=leetcode-75
func uniquePaths(m int, n int) int {
	// return Medium.UniquePaths(m, n)
	return Medium.UniquePathsMemorisation(m, n, make(map[string]int))
}
