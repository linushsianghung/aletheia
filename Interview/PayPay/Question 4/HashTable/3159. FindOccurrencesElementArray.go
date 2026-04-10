package HashTable

// https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/description/
/*
You are given an integer array nums, an integer array queries, and an integer x.

For each queries[i], you need to find the index of the queries[i]th occurrence of x in the nums array. If there are fewer than queries[i] occurrences of x, the answer should be -1 for that query.

Return an integer array answer containing the answers to all queries.
*/
func occurrencesOfElement(nums []int, queries []int, x int) []int {
	count, note := 0, make(map[int]int) // count -> index
	for i, num := range nums {
		if num == x {
			count++
			note[count] = i
		}
	}

	answer := make([]int, len(queries))
	for i, query := range queries {
		if idx, ok := note[query]; ok {
			answer[i] = idx
			continue
		}

		answer[i] = -1
	}

	return answer
}
