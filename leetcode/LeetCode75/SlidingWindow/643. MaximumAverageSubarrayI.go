package SlidingWindow

import "github.com/linushung/aletheia/Concepts/Algorithm/Sliding_Window"

// https://leetcode.com/problems/maximum-average-subarray-i/description/?envId=leetcode-75
func findMaxAverage(nums []int, k int) float64 {
	return Sliding_Window.FindMaxAverage(nums, k)
}
