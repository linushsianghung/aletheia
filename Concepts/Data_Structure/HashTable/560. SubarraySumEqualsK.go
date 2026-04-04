package HashTable

// https://leetcode.com/problems/subarray-sum-equals-k/description/
/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.
*/
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	// Map stores: prefixSum -> frequency
	prefixSumCounts := make(map[int]int)
	prefixSumCounts[0] = 1

	for _, num := range nums {
		sum += num
		// If (sum - k) exists in map, it means we found subarrays summing to k
		count += prefixSumCounts[sum-k]
		prefixSumCounts[sum]++
	}

	return count
}
