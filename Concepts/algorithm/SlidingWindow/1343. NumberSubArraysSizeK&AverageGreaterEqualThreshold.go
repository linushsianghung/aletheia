package SlidingWindow

// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold
/*
Given an array of integers arr and two integers k and threshold, return the number of sub-arrays of size k and average greater than or equal to threshold.
*/
func numOfSubarrays(arr []int, k int, threshold int) int {
	left, sum, result := 0, 0, 0

	for right := 0; right < len(arr); right++ {
		sum += arr[right]
		if right >= k-1 {
			if sum/k >= threshold {
				result++
			}

			sum -= arr[left]
			left++
		}
	}

	return result
}
