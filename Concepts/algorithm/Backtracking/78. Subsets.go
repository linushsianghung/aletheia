package Backtracking

// Subsets https://leetcode.com/problems/subsets/description/
/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Playground: https://go.dev/play/p/-bJ9YARXF0T
*/
func Subsets(nums []int) [][]int {
	return backtrackSubsets(nums)
	// return subsetsBFS(nums)

	// result := make([][]int, 0)
	// subsetsHelper(&result, []int{}, nums, 0)
	// return result

}

func backtrackSubsets(sources []int) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(sources, processor []int, start int)
	localRecursiveFunc = func(sources, processor []int, start int) {
		result = append(result, processor)

		for i := start; i < len(sources); i++ {
			processor = append(processor, sources[i])
			p := make([]int, len(processor))
			copy(p, processor)
			localRecursiveFunc(sources, p, i+1)
			processor = processor[:len(processor)-1]
		}
	}

	localRecursiveFunc(sources, []int{}, 0)
	return result
}

// Reference: https://leetcode.com/problems/subsets/solutions/122645/3ms-easiest-solution-no-backtracking-no-bit-manipulation-no-dfs-no-bullshit/
func subsetsBFS(nums []int) [][]int {
	result := [][]int{make([]int, 0)}

	for _, num := range nums {
		length := len(result)
		for i := 0; i < length; i++ {
			subset := result[i]
			subset = append([]int{num}, subset...)
			result = append(result, subset)
		}
	}

	return result
}

func subsetsHelper(result *[][]int, sources, processor []int, start int) {
	*result = append(*result, processor)

	for i := start; i < len(sources); i++ {
		processor = append(processor, sources[i])
		p := make([]int, len(processor))
		copy(p, processor)
		subsetsHelper(result, sources, p, i+1)
		processor = processor[:len(processor)-1]
	}
}
