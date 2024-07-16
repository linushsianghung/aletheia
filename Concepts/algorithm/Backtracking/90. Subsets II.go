package Backtracking

import "sort"

// https://leetcode.com/problems/subsets-ii/description/
/*
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Analysis:
[1, 3a, 3b, 7]:

[]
[1]
[1] -> [1, 3a]
[1] -> [1, 3a] -> [1, 3a, 3b]
[1] -> [1, 3a] -> [1, 3a, 3b] -> [1, 3a, 3b, 7]
[1] -> [1, 3a] -> [1, 3a, 7]
[1] -> [1, 3b] => Skip
[1] -> [1, 7]
[3a]
[3a] -> [3a, 3b]
[3a] -> [3a, 3b] -> [3a, 3b, 7]
[3a] -> [3a, 7]
[3b] => skip
[7]
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return backtrackSubsetsWithDup(nums)

	// result := make([][]int, 0)
	// subsetsWithDupHelper(&result, make([]int, 0), nums, 0)
	// return result
}

func backtrackSubsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	var localRecursiveFunc func(nums, processor []int, start int)
	localRecursiveFunc = func(nums, processor []int, start int) {
		result = append(result, processor)

		for i := start; i < len(nums); i++ {
			// Constraints: Under the same processed number, skip the element when this element is the same as the previous one
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			processor = append(processor, nums[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(nums, p, i+1)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc([]int{}, nums, 0)
	return result
}

func subsetsWithDupHelper(result *[][]int, nums, processor []int, start int) {
	*result = append(*result, processor)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		processor = append(processor, nums[i])
		p := make([]int, len(processor))
		copy(p, processor)
		subsetsWithDupHelper(result, nums, p, i+1)
		processor = processor[:len(processor)-1]
	}
}
