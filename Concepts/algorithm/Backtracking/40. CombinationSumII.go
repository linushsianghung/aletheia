package Backtracking

import "sort"

// https://leetcode.com/problems/combination-sum-ii/description/
/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return backtrackCombinationsSum2(candidates, target)

	// result := make([][]int, 0)
	// combinationsSum2Helper(&result, candidates, make([]int, 0), target, 0)
	// return result
}

func backtrackCombinationsSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(processor []int, start, remain int)
	localRecursiveFunc = func(processor []int, start, remain int) {
		if remain < 0 {
			return
		}
		if remain == 0 {
			result = append(result, processor)
			return
		}

		for i := start; i < len(candidates); i++ {
			/*
				Constraints: Under the current processed (previous level) number, skip the element when this element is the same as the previous one to prevent duplicated solutions.
				For example, let's say the array is [1,6,1,2,1,6,1] and the sum is 8.
				Whenever we prevent arising of duplicate solutions, we usually first sort the array in this case.
				So our array becomes [1,1,1,1,2,6,6].
				If that condition was not in place, for the combination [1,1,6] we would have got 4C2 * 2C1 = 12 times. ( i.e picking any two 1s out of four possible 1s and one 6 out of two sixes.
				But we do not need [1,1,6] 12 times. not [2,6] 2 times. We just need the solution without any possible duplicate combination. which is [[1,1,6],[2,6]].
			*/
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
