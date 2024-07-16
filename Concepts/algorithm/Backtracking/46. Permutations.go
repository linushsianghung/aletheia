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

func backtrackPermute(nums []int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(nums, processor []int)
	localRecursiveFunc = func(nums, processor []int) {
		if len(processor) == len(nums) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(nums); i++ {
			// Skip current processed number
			if slices.Contains(processor, nums[i]) {
				continue
			}
			processor = append(processor, nums[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(nums, p)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc(nums, []int{})
	return result
}

func permuteHelper(result *[][]int, nums, processor []int) {
	if len(processor) == len(nums) {
		*result = append(*result, processor)
		return
	}

	for i := 0; i < len(nums); i++ {
		if slices.Contains(processor, nums[i]) {
			continue
		}
		processor = append(processor, nums[i])
		p := make([]int, len(processor))
		copy(p, processor)
		permuteHelper(result, nums, p)
		processor = processor[:len(processor)-1]
	}
}
