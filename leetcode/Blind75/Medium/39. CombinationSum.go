package Medium

import "github.com/linushung/aletheia/Concepts/Algorithm/Backtracking"

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	return Backtracking.CombinationSum(candidates, target)
}
