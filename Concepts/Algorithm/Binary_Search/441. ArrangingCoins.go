package Binary_Search

// TODO: Review
// https://leetcode.com/problems/arranging-coins/description/
// Reference: https://leetcode.com/problems/arranging-coins/solutions/1559930/4-java-solution-with-explanations-iterative-binary-search-algebra-math-and-quadratic-math/
/*
You have n coins, and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins.
The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.

Analysis:
- Question: What is the maximum number of complete rows k you can form with n coins?
- Search Space (Possible number of rows k):
    - Minimum rows: 0 or 1.
    - Maximum rows: n. You can't have more rows than you have coins.
    - Search space: [1, n].
- Monotonic "Check" Function: canBuild(k). "Do we have enough coins to build k complete rows?" The number of coins needed for k rows is k * (k + 1) / 2. So, the check is k * (k + 1) / 2 <= n.
    - Monotonicity: If we can build 5 complete rows, can we also build 4 complete rows? Yes.
    - This time, the function is monotonic in the other direction. If check(k) is true, check(k-1) is also true.
    - The results look like: [T, T, ..., T, F, F, ..., F].
- Goal: Find the maximum number of rows, which is the last T.
*/

func arrangeCoins(n int) int {
	// return arrangeCoinsBF(n)
	return arrangeCoinsBS(n)
}

func arrangeCoinsBS(n int) int {
	// Define Search Space
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2

		// Monotonic Function: Formula for counting coins by staircases
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
