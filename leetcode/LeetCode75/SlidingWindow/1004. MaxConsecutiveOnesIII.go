package SlidingWindow

import "github.com/linushung/aletheia/Concepts/Algorithm/Sliding_Window"

// https://leetcode.com/problems/max-consecutive-ones-iii/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/max-consecutive-ones-iii/solutions/247564/java-c-python-sliding-window/comments/326294/

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/
func longestOnes(nums []int, k int) int {
	return Sliding_Window.LongestOnes(nums, k)
}
