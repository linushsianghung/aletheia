package Medium

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions"

// https://leetcode.com/problems/maximum-subarray/
// Ref:
// - Back To Back SWE: [Max Contiguous Subarray Sum - Cubic Time To Kadane's Algorithm](https://www.youtube.com/watch?v=2MmGzdiKR9Y)
// - NeetCode: [Maximum Subarray](https://www.youtube.com/watch?v=5WZl3MMT0Eg)
/*
Given an integer array nums, find the sub-array with the largest sum, and return its sum.

Analysis:
Key Word => Contiguous

1. O(n^3) time complexity: computing all possibilities of sub-array and starting from the head everytime
2. O(n^2) time complexity: using Sliding Window technique when computing all the sub-array
3. O(n^1) time complexity: using Kadane's Algorithm (Dynamic Programming):
	- Inverse the computation to define the sub-problem: what is the maximum sum of contiguous sub-array that end at the index i (ms(i))
	- At each end index of the iteration, what are the choices we have for each element to produce the maximum sum
		- choice 1: ms(i) is index i itself
		- choice 2: ms(i) is the sum of ms(i-1) + i
	- Then max(choice 1, choice 2)
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		/* Kadane's Algorithm (Dynamic Programming):
		   For each element [i], the value of element [i-1] is the best solution of subarray ending at that point.
		   So just consider using the value of element [i] itself or includes the previous best value from element [i-1]

		   currentSum = max(nums[i], currentSum+nums[i])
		*/
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		maxSum = Top_Interview_Questions.Max(maxSum, currentSum)
	}

	return maxSum
}
