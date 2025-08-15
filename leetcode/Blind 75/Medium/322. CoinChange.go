package Medium

import "math"

// https://leetcode.com/problems/coin-change/
// Reference: https://backtobackswe.com/platform/content/the-change-making-problem/solutions
/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an infinite number of each kind of coin.
*/
func coinChange(coins []int, amount int) int {
	return coinChangeDPTabulation(coins, amount)
}

// In order to apply DP Memorisation, below items have to be modified
// 1. "minCount" should be kept in each reminder rather than sharded globally
// 2. Remember the "note" has to be passed through to each stack
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
		// Find the minimum number of coins in each level (sub-problem) by checking the result of using each coin
		for _, coin := range coins {
			count := coinFunc(coins, reminder-coin, note)
			if count >= 0 {
				minCount = min(minCount, count+1)
			}
		}

		// Return -1 to indicate that there is no feasible coin choice for this reminder
		if minCount == math.MaxInt {
			minCount = -1
		}
		note[reminder] = minCount
		// Return the minimum number of coins of this level
		return note[reminder]
	}

	return coinFunc(coins, amount, make(map[int]int))
}

func coinChangeDPMemorisationExercise(coins []int, amount int) int {
	return 0
}

func coinChangeDPTabulation(coins []int, amount int) int {
	// Initialise DP Table
	// The really trick part here is what initial value we can/should ues here... Below values cannot be used because they've been assigned specific meaning already
	// "-1" is defined if that amount of money cannot be made up by any combination of the coins
	// "0" is for the scenario for amount == 0. This is quite weird that is different from "-1"
	// "amount" cannot be used either. It will get problem at scenario coin: [1] and amount = 0
	table := make([]int, amount+1)
	for i := range len(table) {
		// Set default value
		table[i] = amount + 1
	}
	table[0] = 0

	for i := 1; i < len(table); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			table[i] = min(table[i], table[i-coin]+1)
		}
	}
	if table[amount] == amount+1 {
		return -1
	}
	return table[amount]
}

func coinChangeDPTabulationExercise(coins []int, amount int) int {
	return 0
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
