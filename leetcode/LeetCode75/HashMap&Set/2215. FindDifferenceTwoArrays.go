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
	// Use maps as sets for efficient lookups and to handle distinct integers.
	// Using an empty struct `struct{}` is a common Go idiom for sets, as it consumes zero memory, compared to a boolean.
	set1 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	set2 := make(map[int]struct{})
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	// By iterating over the keys of the map `set1`, we are already dealing with unique numbers from `nums1`.
	// This approach avoids iterating through the original arrays again and eliminates the need for slices.Contains.
	// The overall time complexity becomes O(N + M).
	answer1 := make([]int, 0)
	for num := range set1 {
		if _, found := set2[num]; found {
			continue
		}
		answer1 = append(answer1, num)
	}

	// Similarly, for answer2, we iterate over the unique numbers in `set2`  and check for their absence in `set1`.
	answer2 := make([]int, 0)
	for num := range set2 {
		if _, found := set1[num]; found {
			continue
		}
		answer2 = append(answer2, num)
	}

	return [][]int{answer1, answer2}
}

func findDifferenceNaive(nums1 []int, nums2 []int) [][]int {
	note1, note2 := make(map[int]bool), make(map[int]bool)

	for _, num := range nums1 {
		note1[num] = true
	}

	for _, num := range nums2 {
		note2[num] = true
	}

	answer1, answer2 := make([]int, 0), make([]int, 0)
	/*
		It correctly uses maps (note1, note2) for a fast O(1) average time check to see if a number from one array exists in the other.
		However, it also uses slices.Contains(answer1, num) which has to scan the answer1 slice from the beginning each time it's called.

		The time complexity of this part is roughly O(ND1 + MD2), where N and M are the lengths of the arrays, and D1 and D2 are the counts of unique elements.
		This can be quite slow for large inputs.
	*/
	for _, num := range nums1 {
		if ok := note2[num]; ok || slices.Contains(answer1, num) {
			continue
		}
		answer1 = append(answer1, num)
	}
	for _, num := range nums2 {
		if ok := note1[num]; ok || slices.Contains(answer2, num) {
			continue
		}
		answer2 = append(answer2, num)
	}

	return [][]int{answer1, answer2}
}
