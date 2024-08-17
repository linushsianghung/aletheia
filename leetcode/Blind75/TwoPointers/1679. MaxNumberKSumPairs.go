package TwoPointers

import "sort"

// https://leetcode.com/problems/max-number-of-k-sum-pairs/description/?envId=leetcode-75
/*
You are given an integer array nums and an integer k.
In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.
*/
func maxOperations(nums []int, k int) int {
	return maxOperationsMap(nums, k)
}

func maxOperationsMap(nums []int, k int) int {
	count, note := 0, make(map[int]int)

	for _, num := range nums {
		if value := note[k-num]; value > 0 {
			count++
			note[k-num]--
		} else {
			note[num]++
		}
	}

	return count
}

func maxOperations2Pointers(nums []int, k int) int {
	sort.Ints(nums)

	left, right, count := 0, len(nums)-1, 0
	for left < right {
		sum := nums[left] + nums[right]

		if sum < k {
			left++
		} else if sum > k {
			right--
		} else {
			count++
			left++
			right--
		}
	}

	return count
}
