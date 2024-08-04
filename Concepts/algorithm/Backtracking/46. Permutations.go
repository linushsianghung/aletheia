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

func backtrackPermute(source []int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(processor []int)
	localRecursiveFunc = func(processor []int) {
		if len(processor) == len(source) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(source); i++ {
			// Skip current processed number
			if slices.Contains(processor, source[i]) {
				continue
			}
			processor = append(processor, source[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(p)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc([]int{})
	return result
}

func backtrackPermuteExercise(source []int) [][]int {

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
