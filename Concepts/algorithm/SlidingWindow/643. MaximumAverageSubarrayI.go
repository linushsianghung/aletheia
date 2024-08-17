package SlidingWindow

// FindMaxAverage https://leetcode.com/problems/maximum-average-subarray-i/
/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than 10-5 will be accepted.
*/
func FindMaxAverage(nums []int, k int) float64 {
	maxAve := float64(-10_000)
	left, sum := 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if right >= k-1 {
			currentAve := float64(sum) / float64(k)
			if currentAve > maxAve {
				maxAve = currentAve
			}

			sum -= nums[left]
			left++
		}
	}

	return maxAve
}
