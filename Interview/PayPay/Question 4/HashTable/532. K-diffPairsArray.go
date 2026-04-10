package HashTable

// https://leetcode.com/problems/k-diff-pairs-in-an-array/description/
/*
Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.

A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:

0 <= i, j < nums.length
i != j
|nums[i] - nums[j]| == k
Notice that |val| denotes the absolute value of val.
*/
func findPairs(nums []int, k int) int {
	note := make(map[int]int)
	for _, num := range nums {
		note[num]++
	}

	count := 0
	for num, val := range note {
		if k == 0 {
			if val > 1 {
				count++
			}
		} else {
			if _, ok := note[num+k]; ok {
				count++
			}
		}
	}

	return count
}
