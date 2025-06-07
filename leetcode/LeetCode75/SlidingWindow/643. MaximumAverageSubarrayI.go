package SlidingWindow

import "github.com/linushung/aletheia/Concepts/algorithm/SlidingWindow"

// https://leetcode.com/problems/maximum-average-subarray-i/description/?envId=leetcode-75
func findMaxAverage(nums []int, k int) float64 {
	return SlidingWindow.FindMaxAverage(nums, k)
}
