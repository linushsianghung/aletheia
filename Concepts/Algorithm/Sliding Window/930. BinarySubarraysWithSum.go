package Sliding_Window

// https://leetcode.com/problems/binary-subarrays-with-sum/description/
/*
Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

A subarray is a contiguous part of the array.
*/

func numSubarraysWithSum(nums []int, goal int) int {
	// atMost(goal) includes all sets of windows whose total sum is equal to 0 to goal, while atMost(goal-1) comprises sets with sums of 0 to goal-1.
	return numSubarraysWithSumHelper(nums, goal) - numSubarraysWithSumHelper(nums, goal-1)
}

/*
Reference: https://leetcode.com/problems/binary-subarrays-with-sum/editorial/#approach-2-sliding-window
Analysis:
In a standard sliding window approach, once the sum reaches the target goal, the typical strategy involves simply moving the left pointer of the window forward to
potentially find more subarrays. However, this approach has a critical limitation when applied to binary arrays, that is: including a zero element in the subarray won't change the sum.
As a result, even if the sum reaches the goal initially, we might miss further subarrays that also meet the goal by simply shrinking the window as long as the sum
remains equal to the goal. This is because the presence of zeros creates the possibility of combining them with elements encountered later to reach the target sum.
*/
func numSubarraysWithSumHelper(nums []int, goal int) int {
	winStart, sum, count := 0, 0, 0

	for winEnd := range nums {
		sum += nums[winEnd]

		for sum > goal && winStart <= winEnd {
			sum -= nums[winStart]
			winStart++
		}

		count += winEnd - winStart + 1
	}

	return count
}
