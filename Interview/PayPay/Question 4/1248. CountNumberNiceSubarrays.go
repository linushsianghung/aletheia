package Question_4

import "github.com/linushung/aletheia/Concepts/Algorithm/Sliding_Window"

// https://leetcode.com/problems/count-number-of-nice-subarrays/description/
// Reference: https://chatgpt.com/share/69dc3e9e-bfa4-8320-939e-08fe83358f00
/*
Given an array of integers nums and an integer k. A continuous subarray is called nice if there are k odd numbers on it.

Return the number of nice sub-arrays.
*/
func numberOfSubarrays(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
	//return numberOfSubarraysPreSum(nums, k)
}

func atMost(nums []int, k int) int {
	winStart, count, result := 0, 0, 0

	for winEnd, num := range nums {
		if num%2 == 1 {
			count++

			for count > k {
				if nums[winStart]%2 == 1 {
					count--
				}
				winStart++
			}
		}

		result += winEnd - winStart + 1
	}

	return result
}

// Similar solution with 1004. Max Consecutive Ones III: https://leetcode.com/problems/max-consecutive-ones-iii/description/
func longestOnes(nums []int, k int) int {
	return Sliding_Window.LongestOnes(nums, k)
}

func numberOfSubarraysPreSum(nums []int, k int) int {
	count := make(map[int]int)
	count[0] = 1 // base case

	prefixSum := 0
	result := 0

	for _, num := range nums {
		// Convert to 1 (odd) or 0 (even)
		if num%2 == 1 {
			prefixSum++
		}

		// Check how many times (prefixSum - k) occurred
		if val, ok := count[prefixSum-k]; ok {
			result += val
		}

		// Record current prefixSum
		count[prefixSum]++
	}

	return result
}
