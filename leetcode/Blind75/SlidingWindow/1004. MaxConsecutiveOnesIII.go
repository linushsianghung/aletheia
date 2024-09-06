package SlidingWindow

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions"

// https://leetcode.com/problems/max-consecutive-ones-iii/description/?envId=leetcode-75

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/
func longestOnes(nums []int, k int) int {
	left, zeroCount, result := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		result = Top_Interview_Questions.Max(result, right-left+1)
	}

	return result
}
