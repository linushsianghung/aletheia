package HashMap_Set

import "slices"

// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/?envId=leetcode-75
/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	note1, note2 := make(map[int]bool), make(map[int]bool)
	for _, num := range nums1 {
		note1[num] = true
	}
	for _, num := range nums2 {
		note2[num] = true
	}

	answer := make([][]int, 2)
	answer[0] = make([]int, 0)
	answer[1] = make([]int, 0)

	for _, num := range nums1 {
		if ok := note2[num]; ok || slices.Contains(answer[0], num) {
			continue
		}
		answer[0] = append(answer[0], num)
	}
	for _, num := range nums2 {
		if ok := note1[num]; ok || slices.Contains(answer[1], num) {
			continue
		}
		answer[1] = append(answer[1], num)
	}

	return answer
}

func findDifferenceExercise(nums1 []int, nums2 []int) [][]int {
	return nil
}
