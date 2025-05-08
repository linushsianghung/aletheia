package Backtracking

import "sort"

// https://leetcode.com/problems/permutations-ii/description/
// Reference: https://anj910.medium.com/leetcode-47-permutations-ii-7d988b76a1a1
/* Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order. */
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return backtrackPermuteUnique(nums)

	// result := make([][]int, 0)
	// permuteUniqueHelper(&result, nums, make([]int, 0), make([]bool, len(nums)))
	// return result
}

func backtrackPermuteUnique(sources []int) [][]int {
	result := make([][]int, 0)

	var permuteFunc func(processor []int, used []bool)
	permuteFunc = func(processor []int, used []bool) {
		if len(processor) == len(sources) {
			result = append(result, processor)
			return
		}

		for i := 0; i < len(sources); i++ {
			/*
				// Reference: https://leetcode.com/problems/permutations-ii/solutions/18594/really-easy-java-solution-much-easier-than-the-solutions-with-very-high-vote/comments/324818/
				1. The problem is how to handle the duplicates, like [1a, 1b, 2], the results would be [1a, 1b, 2], [1b, 1a, 2]... One way to avoid duplicates is to make sure "1a goes before 1b"
				// Reference: https://leetcode.com/problems/permutations-ii/solutions/18594/really-easy-java-solution-much-easier-than-the-solutions-with-very-high-vote/comments/250112/
				2. Both !use[i - 1] and use[i - 1] are valid, but !use[i - 1] is more efficient.
			*/
			if used[i] || (i > 0 && sources[i-1] == sources[i] && !used[i-1]) {
				continue
			}

			used[i] = true
			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			permuteFunc(p, used)
			processor = processor[:len(processor)-1]
			used[i] = false
		}
	}

	permuteFunc([]int{}, make([]bool, len(sources)))
	return result
}

func backtrackPermuteUniqueExercise(sources []int) [][]int {
	return nil
}

func permuteUniqueHelper(result *[][]int, sources, processor []int, used []bool) {
	if len(processor) == len(sources) {
		*result = append(*result, processor)
		return
	}

	for i := 0; i < len(sources); i++ {
		if used[i] || (i > 0 && sources[i] == sources[i-1] && !used[i-1]) {
			continue
		}

		processor = append(processor, sources[i])
		used[i] = true
		p := make([]int, len(processor))
		copy(p, processor)
		permuteUniqueHelper(result, sources, p, used)
		processor = processor[:len(processor)-1]
		used[i] = false
	}
}
