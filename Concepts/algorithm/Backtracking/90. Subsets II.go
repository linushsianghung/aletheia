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
	// subsetsWithDupHelper(&result, nums, make([]int, 0), 0)
	// return result
}

func backtrackSubsetsWithDup(sources []int) [][]int {
	result := make([][]int, 0)
	var localRecursiveFunc func(processor []int, start int)
	localRecursiveFunc = func(processor []int, start int) {
		result = append(result, processor)

		for i := start; i < len(sources); i++ {
			// Constraints: Under the same processed (previous level) number, skip the element when this element is the same as the previous one
			if i > start && sources[i] == sources[i-1] {
				continue
			}

			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(p, i+1)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc([]int{}, 0)
	return result
}

func subsetsWithDupHelper(result *[][]int, sources, processor []int, start int) {
	*result = append(*result, processor)

	for i := start; i < len(sources); i++ {
		if i > start && sources[i] == sources[i-1] {
			continue
		}

		processor = append(processor, sources[i])
		p := make([]int, len(processor))
		copy(p, processor)
		subsetsWithDupHelper(result, sources, p, i+1)
		processor = processor[:len(processor)-1]
	}
}
