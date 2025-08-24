package SlidingWindow

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/solutions/708112/java-c-python-sliding-window-at-most-one-0
/*
Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

Related Problem: 1004. Max Consecutive Ones III: https://leetcode.com/problems/max-consecutive-ones-iii/description/
*/
func longestSubarray(nums []int) int {
	// Using zeroCount = 0 to be consistent with template
	winStart, zeroCount, maxLen := 0, 0, 0

	for winEnd := range nums {
		if nums[winEnd] == 0 {
			zeroCount++
		}
		// Dynamic-Size Sliding_Window (Double For Loop): skip to the next character of the repeating one
		for zeroCount > 1 {
			if nums[winStart] == 0 {
				zeroCount--
			}

			winStart++
		}

		// Because the element will be deleted, it's unnecessary to plus 1 back
		maxLen = max(maxLen, winEnd-winStart)
	}
	return maxLen
}

func longestSubarrayAlternate(nums []int) int {
	winStart, tune, maxLen := 0, 0, 1

	for winEnd := range nums {
		if nums[winEnd] == 0 {
			tune--
		}

		// Try to find 0 at the most left hand side
		for tune < 0 {
			if nums[winStart] == 0 {
				tune++
			}

			winStart++
		}

		// Because the element will be deleted, it's unnecessary to plus 1 back
		maxLen = max(maxLen, winEnd-winStart)
	}
	return maxLen
}
