package Sliding_Window

// LongestOnes https://leetcode.com/problems/max-consecutive-ones-iii/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/max-consecutive-ones-iii/solutions/247564/java-c-python-sliding-window/comments/326294/
/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/
func LongestOnes(nums []int, k int) int {
	winStart, flip, maxLen := 0, 0, 0

	for winEnd, num := range nums {
		if num == 0 {
			flip++

			for flip > k {
				if nums[winStart] == 0 {
					flip--
				}

				winStart++
			}
		}

		maxLen = max(maxLen, winEnd-winStart+1)
	}

	return maxLen
}

func longestOnesExercise(nums []int, k int) int {
	return 0
}
