package DP_Multidimensional

import "github.com/linushung/aletheia/leetcode/Analysis/Best_Time_Buy_Sell_Stock"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/?envId=leetcode-75

func maxProfitFee(prices []int, fee int) int {
	return Best_Time_Buy_Sell_Stock.MaxProfitFee(prices, fee)
}
