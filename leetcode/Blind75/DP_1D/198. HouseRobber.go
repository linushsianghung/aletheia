package DP_1D

// https://leetcode.com/problems/house-robber/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/house-robber/solutions/156523/from-good-to-great-how-to-approach-most-of-dp-problems
/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected,
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/
func rob(nums []int) int {
	houses := len(nums)
	note := make([]int, houses)
	for i := range note {
		note[i] = -1
	}

	var localFunc func(index int) int
	localFunc = func(index int) int {
		if index < 0 {
			return 0
		}

		if value := note[index]; value >= 0 {
			return value
		}

		note[index] = max(nums[index]+localFunc(index-2), localFunc(index-1))
		return note[index]
	}

	return localFunc(houses - 1)
}
