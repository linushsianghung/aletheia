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
	used := make([]bool, len(sources))

	var permuteFunc func(processor []int)
	permuteFunc = func(processor []int) {
		if len(processor) == len(sources) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(sources); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			permuteFunc(p)
			processor = processor[:len(processor)-1]
			used[i] = false
		}
	}
	permuteFunc(make([]int, 0))

	// Alternative template
	{
		var permuteAltFunc func(processor []int)
		permuteAltFunc = func(processor []int) {
			if len(processor) == len(sources) {
				result = append(result, processor)
				return
			}

			for i := 0; i < len(sources); i++ {
				// Skip current processed number; It might be easier to understand but cannot align to the template with Permutations II
				if slices.Contains(processor, sources[i]) {
					continue
				}
				processor = append(processor, sources[i])
				p := make([]int, len(processor))
				copy(p, processor)
				permuteAltFunc(p)
				processor = processor[:len(processor)-1]
			}
		}
		permuteAltFunc(make([]int, 0))
	}

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
