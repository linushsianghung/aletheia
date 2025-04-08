package BinarySearch

// TODO: Review
// https://leetcode.com/problems/arranging-coins/description/
// Reference: https://leetcode.com/problems/arranging-coins/solutions/1559930/4-java-solution-with-explanations-iterative-binary-search-algebra-math-and-quadratic-math/
/*
You have n coins, and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins.
The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.
*/

func arrangeCoins(n int) int {
	// return arrangeCoinsBF(n)
	return arrangeCoinsBS(n)
}

func arrangeCoinsBS(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2

		// Binary Search Processor: Formula for counting coins by staircases
		coins := (mid + 1) * mid / 2
		if coins > n {
			right = mid - 1
		} else if coins < n {
			left = mid + 1
		} else {
			return mid
		}
	}

	return right
}

func arrangeCoinsBF(n int) int {
	row := 0

	for coin := n; coin > 0; {
		row++
		coin -= row

		if coin == 0 {
			return row
		}
	}

	return row - 1
}
