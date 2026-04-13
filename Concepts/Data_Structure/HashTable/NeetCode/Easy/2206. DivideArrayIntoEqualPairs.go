package Easy

// https://leetcode.com/problems/divide-array-into-equal-pairs/description/
/*
You are given an integer array nums consisting of 2 * n integers.

You need to divide nums into n pairs such that:

Each element belongs to exactly one pair.
The elements present in a pair are equal.
Return true if nums can be divided into n pairs, otherwise return false.
*/
func divideArray(nums []int) bool {
	note := make(map[int]int)

	for _, num := range nums {
		note[num]++
	}

	for _, val := range note {
		if val%2 != 0 {
			return false
		}
	}

	return true
}
