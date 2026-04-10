package Question_4

// ColorTheArray https://leetcode.com/problems/number-of-adjacent-elements-with-the-same-color/description/
/*
You are given an integer n representing an array colors of length n where all elements are set to 0's meaning uncolored. You are also given a 2D integer array queries where queries[i] = [indexi, colori]. For the ith query:

Set colors[indexi] to colori.
Count the number of adjacent pairs in colors which have the same color (regardless of colori).

Return an array answer of the same length as queries where answer[i] is the answer to the ith query.
*/
func ColorTheArray(n int, queries [][]int) []int {
	count, colors := 0, make([]int, n)

	result := make([]int, len(queries))
	for i, q := range queries {
		idx, color := q[0], q[1]

		// Step 1: Remove old contribution
		if colors[idx] != 0 {
			if idx > 0 && colors[idx] == colors[idx-1] {
				count--
			}
			if idx < n-1 && colors[idx] == colors[idx+1] {
				count--
			}
		}

		// Step 2: Apply new color
		colors[idx] = color

		// Step 3: Add new contribution
		if idx > 0 && colors[idx] == colors[idx-1] {
			count++
		}
		if idx < n-1 && colors[idx] == colors[idx+1] {
			count++
		}

		result[i] = count
	}

	return result
}
