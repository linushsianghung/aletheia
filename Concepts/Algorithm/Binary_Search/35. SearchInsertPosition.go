package Binary_Search

// SearchInsert https://leetcode.com/problems/search-insert-position/description/
// Ref: https://leetcode.com/problems/search-insert-position/solutions/249092/come-on-forget-the-binary-search-pattern-template-try-understand-it/?orderBy=most_votes
/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an Algorithm with O(log n) runtime complexity.

Analysis:
Sample1: [1, 3, 5, 8, 10], target=4
left, right := 0, 4: mid = (0 + 4) / 2 = 2 => nums[2] > 4 => right = 2 - 1
left, right := 0, 1: mid = (0 + 1) / 2 = 0 => nums[2] < 4 => left = 0 + 1
left, right := 1, 1: mid = (1 + 1) / 2 = 1 => nums[1] < 4 => left = 1 + 1

left = 2, right = 1

Sample2: [1, 3, 5, 8, 10], target=9
left, right := 0, 4: mid = (0 + 4) / 2 = 2 => nums[2] < 9 => left = 0 + 1
left, right := 1, 4: mid = (1 + 4) / 2 = 2.5 => nums[2] < 9 => left = 1 + 1
left, right := 2, 4: mid = (2 + 4) / 2 = 3 => nums[3] < 9 => left = 2 + 1
left, right := 3, 4: mid = (3 + 4) / 2 = 3.5 => nums[3] < 9 => left = 3 + 1
left, right := 4, 4: mid = (4 + 4) / 2 = 4 => nums[4] > 9 => right = 4 - 1

left = 4, right = 3
*/
func SearchInsert(nums []int, target int) int {
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
