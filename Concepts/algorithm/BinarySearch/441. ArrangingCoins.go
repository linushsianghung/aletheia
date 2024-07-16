package BinarySearch

// https://leetcode.com/problems/arranging-coins/description/
// Reference: https://leetcode.com/problems/arranging-coins/solutions/1559930/4-java-solution-with-explanations-iterative-binary-search-algebra-math-and-quadratic-math/
/*
You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins.
The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.
*/

func arrangeCoins(n int) int {
	// return arrangeCoinsBF(n)
	return arrangeCoinsBS(n)
}

func arrangeCoinsBS(n int) int {

	max := func(x, y int) int {
		if x >= y {
			return x
		} else {
			return y
		}
	}

	left, right := 1, n
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		coins := (mid + 1) * mid / 2

		if coins > n {
			right = mid - 1
		} else {
			left = mid + 1
			result = max(mid, result)
		}
	}

	return result
}

func arrangeCoinsBF(n int) int {
	row := 1

	for coin := n; coin > 0; {
		coin -= row
		if coin == 0 {
			return row
		}
		if coin > 0 {
			row++
		}
	}
	return row - 1
}
