package DP_1D

// Rob https://leetcode.com/problems/house-robber/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/house-robber/solutions/156523/from-good-to-great-how-to-approach-most-of-dp-problems
/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected,
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/
func Rob(nums []int) int {
	note := make([]int, len(nums))
	for i := range note {
		note[i] = -1
	}

	var robFunc func(index int) int
	robFunc = func(index int) int {
		if index < 0 {
			return 0
		}

		if note[index] > -1 {
			return note[index]
		}

		note[index] = max(nums[index]+robFunc(index-2), robFunc(index-1))
		return note[index]
	}

	return robFunc(len(nums) - 1)
}

func robMapNote(nums []int) int {
	note := make(map[int]int)

	var robFunc func(index int) int
	robFunc = func(index int) int {
		if index < 0 {
			return 0
		}

		if val, ok := note[index]; ok {
			return val
		}

		note[index] = max(robFunc(index-1), nums[index]+robFunc(index-2))
		return note[index]
	}

	return robFunc(len(nums) - 1)
}

func robExercise(nums []int) int {
	return 0
}

// Related Problem: 213. House Robber II: https://leetcode.com/problems/house-robber-ii
/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	note := make([]int, len(nums))

	var robFunc func(houses []int, index int) int
	robFunc = func(houses []int, index int) int {
		if index < 0 {
			return 0
		}

		if note[index] > -1 {
			return note[index]
		}

		note[index] = max(houses[index]+robFunc(houses, index-2), robFunc(houses, index-1))
		return note[index]
	}

	for i := range nums {
		note[i] = -1
	}
	result1 := robFunc(nums[:len(nums)-1], len(nums)-2)
	for i := range nums {
		note[i] = -1
	}
	result2 := robFunc(nums[1:], len(nums)-2)

	return max(result1, result2)
}
