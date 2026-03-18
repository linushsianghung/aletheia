package SlidingWindow

import "github.com/linushung/aletheia/Concepts/Algorithm/Sliding_Window"

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/solutions/708112/java-c-python-sliding-window-at-most-one-0
/*
Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

Related Problem: 1004. Max Consecutive Ones III: https://leetcode.com/problems/max-consecutive-ones-iii/description/
*/
func longestSubarray(nums []int) int {
	return Sliding_Window.LongestSubarray(nums)
}
