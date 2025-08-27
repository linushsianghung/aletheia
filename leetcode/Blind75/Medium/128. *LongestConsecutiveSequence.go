package Medium

// https://leetcode.com/problems/longest-consecutive-sequence
// Reference: https://www.youtube.com/watch?v=P6RZZMu_maU
/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

***Failed in specific case: Time Limit Exceeded ***
*/
func longestConsecutive(nums []int) int {
	note := make(map[int]bool)
	for _, num := range nums {
		note[num] = true
	}

	maxLen := 0
	for i := range nums {
		if ok := note[nums[i-1]]; !ok {
			length := 0

			for note[nums[i]] {
				length++
			}

			maxLen = max(maxLen, length)
		}
	}

	return maxLen
}
