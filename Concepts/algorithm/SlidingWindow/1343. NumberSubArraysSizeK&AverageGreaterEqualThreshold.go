package SlidingWindow

// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold
/*
Given an array of integers arr and two integers k and threshold, return the number of sub-arrays of size k and average greater than or equal to threshold.
*/
func numOfSubarrays(arr []int, k int, threshold int) int {
	winStart, sum, count := 0, 0, 0

	for winEnd, num := range arr {
		sum += num
		if winEnd >= k-1 {
			if sum/k >= threshold {
				count++
			}

			sum -= arr[winStart]
			winStart++
		}
	}

	return count
}

func numOfSubarraysExercise(arr []int, k int, threshold int) int {

	return 0
}
