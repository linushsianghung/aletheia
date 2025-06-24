package DP_Multidimensional

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solutions/108871/2-solutions-2-states-dp-solutions-clear-explanation
/*
You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
Find the maximum profit you can achieve. You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.

Note:
- You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
- The transaction fee is only charged once for each stock purchase and sale.
*/

func MaxProfitFee(prices []int, fee int) int {
	days := len(prices)
	buy, sell := make([]int, days), make([]int, days)

	buy[0] = -prices[0]
	for i := 1; i < days; i++ {
		// keep the same as day i-1, or buy from sell status at day i-1
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		// keep the same as day i-1, or sell from buy status at day i-1
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
	}

	return sell[days-1]
}

func maxProfitFeeExercise(prices []int, fee int) int {
	return 0
}
