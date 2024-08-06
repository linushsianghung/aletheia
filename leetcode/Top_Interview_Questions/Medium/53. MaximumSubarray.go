package Medium

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions"

// https://leetcode.com/problems/maximum-subarray/
/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		maxSum = Top_Interview_Questions.Max(maxSum, currentSum)
	}

	return maxSum
}
