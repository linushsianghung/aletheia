package Backtracking

import "slices"

// https://leetcode.com/problems/combination-sum-ii/description/
/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.
*/
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	return backtrackCombinationsSum2(candidates, target)

	// result := make([][]int, 0)
	// combinationsSum2Helper(&result, make([]int, 0), candidates, target, 0)
	// return result
}

func backtrackCombinationsSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(processor, candidates []int, start, remain int)
	localRecursiveFunc = func(processor, candidates []int, start, remain int) {
		if remain < 0 {
			return
		} else if remain == 0 {
			result = append(result, processor)
		} else {
			for i := start; i < len(candidates); i++ {
				// Constraints
				if i > start && candidates[i] == candidates[i-1] {
					continue
				}

				processor = append(processor, candidates[i])
				p := make([]int, len(processor))
				copy(p, processor)
				localRecursiveFunc(candidates, p, i+1, remain-candidates[i])
				processor = processor[:len(processor)-1]
			}
		}
	}

	localRecursiveFunc(candidates, []int{}, target, 0)

	return result
}

func combinationsSum2Helper(result *[][]int, candidates, processor []int, start, remain int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		*result = append(*result, processor)
	} else {
		for i := start; i < len(candidates); i++ {
			// Constraints
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			processor = append(processor, candidates[i])
			p := make([]int, len(processor))
			copy(p, processor)
			combinationsSum2Helper(result, candidates, p, i+1, remain-candidates[i])
			processor = processor[:len(processor)-1]
		}
	}
}
