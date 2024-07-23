package Backtracking

import "slices"

// Permute https://leetcode.com/problems/permutations/description/
/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/
func Permute(nums []int) [][]int {
	slices.Sort(nums)
	return backtrackPermute(nums)

	// result := make([][]int, 0)
	// permuteHelper(&result, make([]int, 0), nums)
	// return result
}

func backtrackPermute(sources []int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(processor []int)
	localRecursiveFunc = func(processor []int) {
		if len(processor) == len(sources) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(sources); i++ {
			// Skip current processed number
			if slices.Contains(processor, sources[i]) {
				continue
			}
			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(p)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc([]int{})
	return result
}

func permuteHelper(result *[][]int, sources, processor []int) {
	if len(processor) == len(sources) {
		*result = append(*result, processor)
		return
	}

	for i := 0; i < len(sources); i++ {
		if slices.Contains(processor, sources[i]) {
			continue
		}
		processor = append(processor, sources[i])
		p := make([]int, len(processor))
		copy(p, processor)
		permuteHelper(result, sources, p)
		processor = processor[:len(processor)-1]
	}
}
