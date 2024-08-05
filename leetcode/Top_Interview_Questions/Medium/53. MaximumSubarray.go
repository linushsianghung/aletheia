package Medium

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

		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func maxF(x, y int) int {
	if x > y {
		return x
	}

	return y
}
