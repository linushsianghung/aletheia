package Sliding_Window

// FindMaxAverage https://leetcode.com/problems/maximum-average-subarray-i/
/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than 10-5 will be accepted.
*/
func FindMaxAverage(nums []int, k int) float64 {
	winStart, maxAvg := 0, float64(-10_000)
	// Window State: sum of subarray for calculating average
	sum := 0

	for winEnd, num := range nums {
		sum += num

		if winEnd >= k-1 {
			maxAvg = max(maxAvg, float64(sum)/float64(k))

			sum -= nums[winStart]
			winStart++
		}
	}

	return maxAvg
}

func FindMaxAverageExercise(nums []int, k int) float64 {
	winStart, maxAvg := 0, float64(10_000)
	sum := 0

	for winEnd, num := range nums {
		sum += num

		if winEnd >= k-1 {
			maxAvg = max(maxAvg, float64(sum)/float64(k))

			sum -= nums[winStart]
			winStart++
		}
	}

	return 0
}
