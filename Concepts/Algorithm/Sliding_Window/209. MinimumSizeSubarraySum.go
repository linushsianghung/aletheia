package Sliding_Window

// https://leetcode.com/problems/minimum-size-subarray-sum/description/
/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target.

If there is no such subarray, return 0 instead.
*/
func minSubArrayLen(target int, nums []int) int {
	winStart, sum, minLen := 0, 0, len(nums)+1

	for winEnd, num := range nums {
		sum += num

		// Dynamic-Size Sliding Window: based on whether sum is larger than target
		for sum >= target {
			minLen = min(minLen, winEnd-winStart+1)

			sum -= nums[winStart]
			winStart++
		}
	}

	if minLen == len(nums)+1 {
		minLen = 0
	}

	return minLen
}

func minSubArrayLenExercise(target int, nums []int) int {
	return 0
}
