package Array

import (
	TwoPointers "github.com/linushung/aletheia/leetcode/Basic/2Pointers"
)

// https://leetcode.com/problems/two-sum/
// Ref: https://leetcode.com/problems/two-sum/solutions/737092/sum-megapost-python3-solution-with-a-detailed-explanation/
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int {
	return mapImpl(nums, target)
}

/*
Instead of using 2 loops to iterate the whole array (brute forcefully), using map to collect each element
and check the difference between current number and target (target-n) already in the map on the way to shortage the
calculation.
*/
func mapImpl(nums []int, target int) []int {
	note := make(map[int]int) // number, index

	for i, num := range nums {
		if j, ok := note[target-num]; ok {
			return []int{i, j}
		}

		note[num] = i
	}

	return nil
}

func mapImplExercise(nums []int, target int) []int {

	return nil
}

/* if the array is sorted => O(n) */
func towPointerImpl(nums []int, target int) []int {
	return TwoPointers.TwoSumII(nums, target)
}

/* if the array is sorted => O(nlogn) */
func binarySearchImpl(nums []int, target int) []int {
	return nil
}

/* O(n^2) */
func bruteForceImpl(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
