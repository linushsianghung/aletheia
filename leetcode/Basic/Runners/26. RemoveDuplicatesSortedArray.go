package Runners

// RemoveDuplicates https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Ref: https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums. Consider the number of
unique elements of nums to be k, to get accepted, you need to do the following things:
- Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
The remaining elements of nums are not important as well as the size of nums.
- Return k.

Analysis:
Sample1: {1, 2, 3, 5, 6, 9}
Sample2: {1, 2, 2, 2, 6, 9}

Keep slow index to point to the duplicated element
*/
func RemoveDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	// Return k, the number of unique elements of nums, namely index + 1
	return slow + 1
}

func removeDuplicatesExercise(nums []int) int {

	return 0
}
