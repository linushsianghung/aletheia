package Backtracking

// https://leetcode.com/problems/combination-sum/description/
/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/
func combinationSum(candidates []int, target int) [][]int {
	// Reference: Sorting Strings In Go(https://words.aead.dev/sortstr.html)
	return backtrackCombinationsSum(candidates, target)

	// result := make([][]int, 0)
	// combinationsSumHelper(&result, candidates, make([]int, 0), target, 0)
	// return result
}

func backtrackCombinationsSum(source []int, target int) [][]int {
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

		for i := start; i < len(source); i++ {
			processor = append(processor, source[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(p, remain-source[i], i)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc([]int{}, target, 0)
	return result
}

func backtrackCombinationsSumExercise(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var localFunc func(processor []int, remain, start int)
	localFunc = func(processor []int, remain, start int) {
		if remain < 0 {
			return
		}

		if remain == 0 {
			result = append(result, processor)
		} else {
			for i := start; i < len(candidates); i++ {
				processor = append(processor, candidates[i])
				p := make([]int, len(processor))
				copy(p, processor)
				localFunc(p, remain-candidates[i], i)
				processor = processor[:len(processor)-1]
			}
		}

	}

	localFunc(make([]int, 0), target, 0)
	return nil
}

func combinationsSumHelper(result *[][]int, sources, processor []int, remain, start int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		*result = append(*result, processor)
	} else {
		for i := start; i < len(sources); i++ {
			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			combinationsSumHelper(result, sources, p, remain-sources[i], i)
			processor = processor[:len(processor)-1]
		}
	}
}
