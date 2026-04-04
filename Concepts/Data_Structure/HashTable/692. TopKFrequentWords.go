package HashTable

import "sort"

// TopKFrequent https://leetcode.com/problems/top-k-frequent-words/description/
/*
Given an array of strings words and an integer k, return the k most frequent strings.

Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.
*/
func TopKFrequent(words []string, k int) []string {
	note := make(map[string]int)
	for _, word := range words {
		note[word]++
	}

	frequency := make([][]string, len(words)+1)
	for key, value := range note {
		if frequency[value] == nil {
			frequency[value] = make([]string, 0)
		}

		frequency[value] = append(frequency[value], key)
	}

	count, result := 0, make([]string, 0)
	for i := len(frequency) - 1; i >= 0; i-- {
		if count >= k {
			break
		}

		sort.Slice(frequency[i], func(x, y int) bool {
			return frequency[i][x] < frequency[i][y]
		})
		result = append(result, frequency[i]...)
		count += len(frequency[i])
	}

	return result[:k]
}
