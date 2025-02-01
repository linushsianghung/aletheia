package SlidingWindow

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/solutions/708112/java-c-python-sliding-window-at-most-one-0/?envType=study-plan-v2&envId=leetcode-75
/*
Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

Related Topic: 1004. Max Consecutive Ones III: https://leetcode.com/problems/max-consecutive-ones-iii/description/
*/
func longestSubarray(nums []int) int {
	// Using buffer = 0 to be consistent with template
	winStart, buffer, maxLen := 0, 0, 0

	for winEnd := 0; winEnd < len(nums); winEnd++ {
		if nums[winEnd] == 0 {
			buffer++
		}
		// Dynamic-Size Sliding Window (Double For Loop): skip to the next character of the repeating one
		for buffer > 1 {
			if nums[winStart] == 0 {
				buffer--
			}

			winStart++
		}

		// Because the element will be deleted, it's unnecessary to plus 1 back
		if winEnd-winStart > maxLen {
			maxLen = winEnd - winStart
		}
	}
	return maxLen
}

func longestSubarrayAlternate(nums []int) int {
	winStart, tune, maxLen := 0, 0, 1

	for winEnd := 0; winEnd < len(nums); winEnd++ {
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
		length := winEnd - winStart
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
