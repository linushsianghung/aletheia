package Easy

// https://leetcode.com/problems/number-of-good-pairs/description/
/*
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/
func numIdenticalPairs(nums []int) int {
	count, note := 0, make(map[int]int) // map { 1: true }
	for _, num := range nums {
		if val, ok := note[num]; ok {
			count += val
		}

		note[num]++
	}

	return count
}
