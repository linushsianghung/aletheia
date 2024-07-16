package BinarySearch

// Search https://leetcode.com/problems/binary-search/
// Ref:
// - https://leetcode.com/problems/binary-search/solutions/423162/binary-search-101/
// - https://leetcode.com/discuss/study-guide/1233854/a-noobs-guide-to-the-binary-search-algorithm
/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.
*/
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	return searchRecursively(nums, left, right, target)
}

func searchRecursively(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}

	// calculate mid := left + (right-left)/2 rather than (right-left)/2 to prevent overflow.
	mid := left + (right-left)/2
	if nums[mid] < target {
		return searchRecursively(nums, mid+1, right, target)
	} else if nums[mid] > target {
		return searchRecursively(nums, left, mid-1, target)
	} else {
		return mid
	}
}

func searchRecursivelyExercise(nums []int, left int, right int, target int) int {
	return 0
}

func searchIteratively(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func searchIterativelyExercise(nums []int, target int) int {
	return 0
}
