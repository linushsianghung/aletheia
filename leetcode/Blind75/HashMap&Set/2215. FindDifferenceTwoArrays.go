package HashMap_Set

// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/?envId=leetcode-75
/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	map1, map2 := make(map[int]bool), make(map[int]bool)

	for _, num := range nums1 {
		map1[num] = true
	}
	for _, num := range nums2 {
		map2[num] = true
	}

	result := make([][]int, 2)
	result[0] = make([]int, 0)
	result[1] = make([]int, 0)
	for num := range map1 {
		if _, ok := map2[num]; ok {
			continue
		}
		result[0] = append(result[0], num)
	}
	for num := range map2 {
		if _, ok := map1[num]; ok {
			continue
		}
		result[1] = append(result[1], num)
	}

	return result
}
