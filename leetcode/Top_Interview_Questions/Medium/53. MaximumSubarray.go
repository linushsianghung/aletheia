package Medium

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions"

// https://leetcode.com/problems/maximum-subarray/
// Ref:
// - Back To Back SWE: [Max Contiguous Subarray Sum - Cubic Time To Kadane's Algorithm](https://www.youtube.com/watch?v=2MmGzdiKR9Y)
// - NeetCode: [Maximum Subarray](https://www.youtube.com/watch?v=5WZl3MMT0Eg)
/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		/* Kadane's Algorithm (Dynamic Programming):
		   For each element [i], the value of element [i-1] is the best solution of subarry ending at that point.
		   So just consider using the value of element [i] itself or includes the previous best value from element [i-1]

		   currentSum = max(nums[i], currentSum+nums[i])
		*/
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		maxSum = Top_Interview_Questions.Max(maxSum, currentSum)
	}

	return maxSum
}
