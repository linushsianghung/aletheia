package Backtracking

import "slices"

// https://leetcode.com/problems/permutations-ii/description/
/* Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order. */
func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	return backtrackPermuteUnique(nums)

	// result := make([][]int, 0)
	// permuteUniqueHelper(&result, make([]int, 0), nums, make([]bool, len(nums)))
	// return result
}

func backtrackPermuteUnique(nums []int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(nums, processor []int, used []bool)
	localRecursiveFunc = func(nums, processor []int, used []bool) {
		if len(processor) == len(nums) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(nums); i++ {
			// TODO: Cannot understand the meaning of this condition
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			processor = append(processor, nums[i])
			used[i] = true
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(nums, p, used)
			processor = processor[:len(processor)-1]
			used[i] = false
		}
	}

	localRecursiveFunc(nums, []int{}, make([]bool, len(nums)))
	return result
}

func permuteUniqueHelper(result *[][]int, nums, processor []int, used []bool) {
	if len(processor) == len(nums) {
		*result = append(*result, processor)
		return
	}

	for i := 0; i < len(nums); i++ {
		// TODO: Cannot understand the meaning of this condition
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}

		processor = append(processor, nums[i])
		used[i] = true
		p := make([]int, len(processor))
		copy(p, processor)
		permuteUniqueHelper(result, nums, p, used)
		processor = processor[:len(processor)-1]
		used[i] = false
	}
}
