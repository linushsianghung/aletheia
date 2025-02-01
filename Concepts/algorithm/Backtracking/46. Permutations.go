package Backtracking

import "slices"

// Permute https://leetcode.com/problems/permutations/description/
/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/
func Permute(nums []int) [][]int {
	return backtrackPermute(nums)

	// result := make([][]int, 0)
	// permuteHelper(&result, nums, make([]int, 0))
	// return result
}

func backtrackPermute(sources []int) [][]int {
	result := make([][]int, 0)

	var nestedFunc func(processor []int, used []bool)
	nestedFunc = func(processor []int, used []bool) {

		if len(processor) == len(sources) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(sources); i++ {
			if used[i] {
				continue
			}

			processor = append(processor, sources[i])
			used[i] = true
			p := make([]int, len(processor))
			copy(p, processor)
			nestedFunc(p, used)
			processor = processor[:len(processor)-1]
			used[i] = false
		}
	}

	// It might be easier to understand but cannot align to the template with Permutations II
	var nestedAltFunc func(processor []int)
	nestedAltFunc = func(processor []int) {
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
			nestedAltFunc(p)
			processor = processor[:len(processor)-1]
		}
	}

	nestedFunc(make([]int, 0), make([]bool, len(sources)))
	return result
}

func backtrackPermuteExercise(sources []int) [][]int {
	return nil
}

func permuteHelper(result *[][]int, source, processor []int) {
	if len(processor) == len(source) {
		*result = append(*result, processor)
		return
	}

	for i := 0; i < len(source); i++ {
		if slices.Contains(processor, source[i]) {
			continue
		}
		processor = append(processor, source[i])
		p := make([]int, len(processor))
		copy(p, processor)
		permuteHelper(result, source, p)
		processor = processor[:len(processor)-1]
	}
}
