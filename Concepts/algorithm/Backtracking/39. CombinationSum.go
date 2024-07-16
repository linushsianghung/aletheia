package Backtracking

import "slices"

// https://leetcode.com/problems/combination-sum/description/
/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/
func combinationSum(candidates []int, target int) [][]int {
	// Reference: Sorting Strings In Go(https://words.aead.dev/sortstr.html)
	slices.Sort(candidates)
	return backtrackCombinationsSum(candidates, target)

	// result := make([][]int, 0)
	// combinationsSumHelper(&result, make([]int, 0), candidates, target, 0)
	// return result
}

func backtrackCombinationsSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(candidates, processor []int, start, remain int)
	localRecursiveFunc = func(candidates, processor []int, start, remain int) {
		if remain < 0 {
			return
		} else if remain == 0 {
			result = append(result, processor)
		} else {
			for i := start; i < len(candidates); i++ {
				processor = append(processor, candidates[i])
				p := make([]int, len(processor))
				copy(p, processor)
				// the "start" is still i because it's allowed to use the same element as many as we want
				localRecursiveFunc(candidates, p, remain-candidates[i], i)
				processor = processor[:len(processor)-1]
			}
		}
	}

	localRecursiveFunc(candidates, []int{}, 0, target)
	return result
}

func combinationsSumHelper(result *[][]int, candidates, processor []int, start, remain int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		*result = append(*result, processor)
	} else {
		for i := start; i < len(candidates); i++ {
			processor = append(processor, candidates[i])
			p := make([]int, len(processor))
			copy(p, processor)
			combinationsSumHelper(result, candidates, p, i, remain-candidates[i])
			processor = processor[:len(processor)-1]
		}
	}
}
