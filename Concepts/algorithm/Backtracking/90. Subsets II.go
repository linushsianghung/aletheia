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
	var nestedFunc func(processor []int, start int)
	nestedFunc = func(processor []int, start int) {
		result = append(result, processor)
		//if len(processor) == len(sources) {
		//	return
		//}

		for i := start; i < len(sources); i++ {
			// Constraints: Skip the element when this is the same as the previous one in each backtrack level in order to prevent duplicate subsets,
			if i > start && sources[i] == sources[i-1] {
				continue
			}

			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			nestedFunc(p, i+1)
			processor = processor[:len(processor)-1]
		}
	}

	nestedFunc([]int{}, 0)
	return result
}

func backtrackSubsetsWithDupExercise(sources []int) [][]int {
	return nil
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
