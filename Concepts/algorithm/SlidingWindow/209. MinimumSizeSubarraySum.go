package SlidingWindow

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum/description/
/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target.

If there is no such subarray, return 0 instead.
*/
func minSubArrayLen(target int, nums []int) int {
	winStart, sum, minLen := 0, 0, math.MaxInt

	for winEnd := 0; winEnd < len(nums); winEnd++ {
		sum += nums[winEnd]
		// Dynamic-Size Sliding Window (Double For Loop)
		for sum >= target {
			currrentLen := winEnd - winStart + 1
			if currrentLen < minLen {
				minLen = currrentLen
			}

			sum -= nums[winStart]
			winStart++
		}
	}

	if minLen == math.MaxInt {
		minLen = 0
	}

	return minLen
}
