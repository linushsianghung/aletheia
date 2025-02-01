package DynamicProgramming

// https://leetcode.com/problems/longest-increasing-subsequence/description/
// Ref: https://www.youtube.com/watch?v=fV-TF4OvZpk
/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	table := make([]int, len(nums))
	for i := range table {
		table[i] = 1
	}

	var maxLen int
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				table[i] = max(table[j]+1, table[i])
			}
		}

		maxLen = max(maxLen, table[i])
	}

	return maxLen
}
