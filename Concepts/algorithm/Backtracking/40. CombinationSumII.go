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
	// combinationsSum2Helper(&result, candidates, make([]int, 0), target, 0)
	// return result
}

func backtrackCombinationsSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(processor []int, remain, start int)
	localRecursiveFunc = func(processor []int, remain, start int) {
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
				localRecursiveFunc(p, remain-candidates[i], i+1)
				processor = processor[:len(processor)-1]
			}
		}
	}

	localRecursiveFunc([]int{}, target, 0)

	return result
}

func combinationsSum2Helper(result *[][]int, sources, processor []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		*result = append(*result, processor)
	} else {
		for i := start; i < len(sources); i++ {
			// Constraints
			if i > start && sources[i] == sources[i-1] {
				continue
			}

			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			combinationsSum2Helper(result, sources, p, remain-sources[i], i+1)
			processor = processor[:len(processor)-1]
		}
	}
}
