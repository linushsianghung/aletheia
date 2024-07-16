package Dynamic_Programming

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description
/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func maxProfit(prices []int) int {

	// return maxProfitSimpleThrough(prices)
	return maxProfit2Pointer(prices)
}

// Ref: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/1735493/java-c-best-ever-explanation-could-possible/?envType=featured-list&envId=top-interview-questions?envType=featured-list&envId=top-interview-questions
func maxProfitSimpleThrough(prices []int) int {
	// Use 2 variables to maintain the lowest price and maximum profit
	lowest, profit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		// Find out the lowest price
		if lowest > prices[i] {
			lowest = prices[i]
		}

		// Then check current profit
		if prices[i]-lowest > profit {
			profit = prices[i] - lowest
		}
	}

	return profit
}

func maxProfitExercise(prices []int) int {
	return 0
}

// Ref: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/1735550/python-javascript-easy-solution-with-very-clear-explanation/
func maxProfit2Pointer(prices []int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	buy, profit := 0, 0
	for sell := 1; sell < len(prices); sell++ {
		currentProfit := prices[sell] - prices[buy]
		if currentProfit > 0 {
			profit = max(currentProfit, profit)
		} else {
			buy = sell
		}
	}

	return profit
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/39038/kadane-s-algorithm-since-no-one-has-mentioned-about-this-so-far-in-case-if-interviewer-twists-the-input/
func maxProfitKadane(prices []int) int {
	return 0
}
