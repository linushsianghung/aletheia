package HashTable

import Question_4 "github.com/linushung/aletheia/Interview/PayPay/Question 4"

// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/description/
/*
You are given an integer limit and a 2D array queries of size n x 2.

There are limit + 1 balls with distinct labels in the range [0, limit]. Initially, all balls are uncolored. For every query in queries that is of the form [x, y], you mark ball x with the color y. After each query, you need to find the number of colors among the balls.

Return an array result of length n, where result[i] denotes the number of colors after ith query.

Note that when answering a query, lack of a color will not be considered as a color.
*/
func queryResults(limit int, queries [][]int) []int {
	ballColor := make(map[int]int)  // ball -> color
	colorCount := make(map[int]int) // color -> frequency

	result := make([]int, len(queries))
	for i, q := range queries {
		idx, color := q[0], q[1]

		// Step 1: remove old color count if exists
		if oldColor, ok := ballColor[idx]; ok {
			colorCount[oldColor]--
			if colorCount[oldColor] == 0 {
				delete(colorCount, oldColor)
			}
		}

		// Step 2: Apply new color
		ballColor[idx] = color

		// Step 3: Add new color count
		colorCount[color]++

		result[i] = len(colorCount)
	}

	return result
}

// Similar solution with 2672. Number of Adjacent Elements With the Same Color: https://leetcode.com/problems/number-of-adjacent-elements-with-the-same-color/description/
func colorTheArray(n int, queries [][]int) []int {
	return Question_4.ColorTheArray(n, queries)
}
