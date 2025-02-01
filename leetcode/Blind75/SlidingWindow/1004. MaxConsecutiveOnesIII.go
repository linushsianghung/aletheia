package SlidingWindow

// https://leetcode.com/problems/max-consecutive-ones-iii/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/max-consecutive-ones-iii/solutions/247564/java-c-python-sliding-window/

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/
func longestOnes(nums []int, k int) int {
	winStart, zeroCount, maxNum := 0, 0, 0

	for winEnd, num := range nums {
		if num == 0 {
			zeroCount++
		}
		// Dynamic-Size Sliding Window (Double For Loop): Try to find 0 at the most left hand side
		for zeroCount > k {
			if nums[winStart] == 0 {
				zeroCount--
			}

			winStart++
		}

		if maxNum < winEnd-winStart+1 {
			maxNum = winEnd - winStart + 1
		}
	}

	return maxNum
}
