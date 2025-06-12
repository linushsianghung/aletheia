package DynamicProgramming

import "math"

// https://leetcode.com/problems/coin-change/description/
// Reference: https://backtobackswe.com/platform/content/the-change-making-problem/solutions
/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an infinite number of each kind of coin.
*/
func coinChange(coins []int, amount int) int {
	return 0
}

// In order to apply DP Memorisation, below items have to be modified
// 1. minCount should be kept in each reminder rather than sharded globally
// 2. Remember the note has to be passed all ways down to each stack
func coinChangeDPMemorisation(coins []int, amount int) int {
	// It's for handling the special cases like coins = [1] & amount = 0
	if amount < 1 {
		return 0
	}

	var coinFunc func(coins []int, reminder int, note map[int]int) int
	coinFunc = func(coins []int, reminder int, note map[int]int) int {
		if reminder < 0 {
			return -1
		}
		if reminder == 0 {
			return 0
		}
		if val, ok := note[reminder]; ok {
			return val
		}

		minCount := math.MaxInt
		for _, coin := range coins {
			result := coinFunc(coins, reminder-coin, note)
			if result >= 0 {
				minCount = min(minCount, result+1)
			}
		}

		// It's for handling the special cases like coins = [2] & amount = 3
		if minCount == math.MaxInt {
			minCount = -1
		}
		note[reminder] = minCount

		return note[reminder]
	}

	result := coinFunc(coins, amount, make(map[int]int))
	return result
}

// It's a straight forward but naive solution because it cannot be easy to apply DP Memorisation afterward.
func coinChangeRecursively(coins []int, amount int) int {
	minCount := math.MaxInt

	var coinFunc func(coins []int, reminder, count int)
	coinFunc = func(coins []int, reminder, count int) {
		if reminder < 0 {
			return
		}
		if reminder == 0 {
			minCount = min(minCount, count)
			return
		}

		for _, coin := range coins {
			coinFunc(coins, reminder-coin, count+1)
		}
	}

	coinFunc(coins, amount, 0)
	if minCount == math.MaxInt {
		minCount = -1
	}
	return minCount
}
