package BinarySearch

// https://leetcode.com/problems/search-insert-position/
// Ref: https://leetcode.com/problems/search-insert-position/solutions/249092/come-on-forget-the-binary-search-pattern-template-try-understand-it/?orderBy=most_votes
/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/
/* Analysis:
Sample: {1, 2, 4, 6, 9, 10, 14, 17, 19, 23, 26}
left = 0 & right = 10
target = 11

*/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}

func searchInsertExercise(nums []int, target int) int {
	return 0
}

func searchInsertAlt(nums []int, target int) int {
	// Set possible index range(inclusive)
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	// 1 element left at the end => post-processing
	if nums[left] < target {
		return left + 1
	} else {
		return left
	}
}
