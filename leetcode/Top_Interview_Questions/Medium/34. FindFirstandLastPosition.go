package Medium

// SearchRange https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description
// Ref: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/solutions/4147878/video-give-me-10-minutes-binary-search-solution-python-javascript-java-c/
/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Analysis: Using Binary Search twice for left most value and right most value
*/
func SearchRange(nums []int, target int) []int {
	result := make([]int, 2)
	result[0] = binarySearchLeft(nums, target)
	result[1] = binarySearchRight(nums, target)

	return result
}

func binarySearchLeft(nums []int, target int) int {
	left, right, index := 0, len(nums)-1, -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			right = mid - 1
		}
	}

	return index
}

func binarySearchLeftExercise(nums []int, target int) int {

	return 0
}

func binarySearchRight(nums []int, target int) int {
	left, right, index := 0, len(nums)-1, -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			left = left + 1
		}
	}

	return index
}

func binarySearchRightExercise(nums []int, target int) int {

	return 0
}
