package Basic

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii
/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.
*/

func maxProfitII(prices []int) int {
	slow, profit := 0, 0

	for fast := 1; fast < len(prices); fast++ {
		if prices[fast-1] > prices[fast] {
			profit += prices[fast-1] - prices[slow]
			slow = fast
		}

		if fast+1 == len(prices) {
			profit += prices[fast] - prices[slow]
		}
	}

	return profit
}
