package HashTable

import (
	"sort"
)

// TopKFrequent https://leetcode.com/problems/top-k-frequent-words/description/
/*
Given an array of strings words and an integer k, return the k most frequent strings.

Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.
*/
func TopKFrequent(words []string, k int) []string {
	note := make(map[string]int) // map: {"i": 2, "love": 2, ...}
	for _, word := range words {
		note[word]++
	}

	frequency := make([][]string, len(words)+1) // slice: [[], ["leetcode", "coding"], ["i", "love"]]
	for word, count := range note {
		if frequency[count] == nil {
			frequency[count] = make([]string, 0)
		}

		frequency[count] = append(frequency[count], word)
	}

	result := make([]string, 0)
	for i := len(frequency) - 1; i >= 0; i-- {
		sort.Slice(frequency[i], func(x, y int) bool {
			return frequency[i][x] < frequency[i][y]
		})

		result = append(result, frequency[i]...)
		if len(result) >= k {
			break
		}
	}

	return result[:k]
}
