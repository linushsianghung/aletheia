package SlidingWindow

// https://leetcode.com/problems/maximum-average-subarray-i/
/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/
func findMaxAverage(nums []int, k int) float64 {
	var max float64
	left, sum := 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if right >= k-1 {
			ave := float64(sum) / float64(k)
			if ave > max {
				max = ave
			}

			sum -= nums[left]
			left++
		}
	}

	return max
}
