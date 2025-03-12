package Backtracking

// https://leetcode.com/problems/combination-sum-iii/description/?envId=leetcode-75
/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
- Only numbers 1 through 9 are used.
- Each number is used at most once.

Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.
*/
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)

	var localFunc func(processor []int, remain, start int)
	localFunc = func(processor []int, remain, start int) {
		if len(processor) > k {
			return
		}

		if len(processor) == k && remain == 0 {
			result = append(result, processor)
			return
		}

		for i := start; i <= 9; i++ {
			processor = append(processor, i)
			p := make([]int, len(processor))
			copy(p, processor)
			localFunc(p, remain-i, i+1)
			processor = processor[:len(processor)-1]
		}
	}

	localFunc(make([]int, 0), n, 1)
	return result
}

func backtrackCombinationSum3Exercise(k int, n int) [][]int {
	return nil
}
