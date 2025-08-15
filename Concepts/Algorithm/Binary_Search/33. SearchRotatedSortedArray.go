package Binary_Search

import "math"

// SearchRotate https://leetcode.com/problems/search-in-rotated-sorted-array/description/
// Ref: https://leetcode.com/problems/search-in-rotated-sorted-array/solutions/154836/The-INF-and-INF-method-but-with-a-better-explanation-for-dummies-like-me/
/*
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an Algorithm with O(log n) runtime complexity.

Big Hint: "If nums[mid] and target are "on the same side" of nums[0], just keep going to normal Binary_Search".
*/
func SearchRotate(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// If both target and nums[mid] are on the same side, just using Binary_Search as usual
		if (target >= nums[0] && nums[mid] >= nums[0]) || (target < nums[0] && nums[mid] < nums[0]) {
			if nums[mid] > target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				return mid
			}
			// If target and nums[mid] are on the different side
		} else {
			// and if target is on the right side of the 0 ([3,4,5,6,7,0,1,2] and target = 1), we should move "left pointer" instead to reduce the search space
			if nums[mid] > target {
				left = mid + 1
				// and if target is on the left side of the 0 ([6,7,0,1,2,3,4,5] and target = 7), we should move "right pointer" instead to reduce the search space
			} else if nums[mid] < target {
				right = mid - 1
			} else {
				return mid
			}
		}
	}

	return -1
}

func SearchRotateExercise(nums []int, target int) int {
	return 0
}

func SearchRotateComparator(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		var comparator int

		// Checking if both target and nums[mid] are on the same side.
		if (target < nums[0] && nums[mid] < nums[0]) || (target >= nums[0] && nums[mid] >= nums[0]) {
			comparator = nums[mid]
		} else {
			// Trying to figure out where nums[mid] is and making comparator as -INF or INF
			if target < nums[0] {
				// target on the right side of 0 so set comparator as math.MinInt to make it smaller than target
				comparator = math.MinInt
			} else {
				// target on the left side of 0 so set comparator as math.MaxInt to make it larger than target
				comparator = math.MaxInt
			}
		}

		if comparator > target {
			right = mid - 1
		} else if comparator < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
