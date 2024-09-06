package Medium

import "math"

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
// Ref: https://leetcode.com/problems/search-in-rotated-sorted-array/solutions/154836/The-INF-and-INF-method-but-with-a-better-explanation-for-dummies-like-me/
/*
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Big Hint: "If nums[mid] and target are "on the same side" of nums[0], just keep going to normal Binary Search".
*/
func search(nums []int, target int) int {
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
				comparator = math.MinInt
			} else {
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

func searchAlt(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// Here I am using more straightforward conditions checking (if target and mid at different side, modify the element to Max or Min, otherwise keep going)
		// The downside is that it seems not dealing with the edge cases that well.
		// I need extra conditions checking as below to satisfy the requirements which makes the solution not quite easy to understand
		// => Add "else { return 0 }" conditions checking when modifying the left or right edge
		if (target > nums[0] && nums[mid] < nums[0]) || (target <= nums[0] && nums[mid] >= nums[0]) {
			if target > nums[0] {
				nums[mid] = math.MaxInt
			} else if target < nums[0] {
				nums[mid] = math.MinInt
			} else {
				return 0
			}
		}

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
