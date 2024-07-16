package TwoPointers

// https://leetcode.com/problems/squares-of-a-sorted-array/
/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
*/
func sortedSquares(nums []int) []int {
	left, right, tail := 0, len(nums)-1, len(nums)-1
	result := make([]int, len(nums))

	for left <= right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			result[tail] = nums[left] * nums[left]
			left++
		} else {
			result[tail] = nums[right] * nums[right]
			right--
		}
		tail--
	}
	return result
}

func sortedSquaresExercise(nums []int) []int {

	return nil
}
