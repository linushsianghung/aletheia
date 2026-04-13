package HashTable

// TopKFrequentElements https://leetcode.com/problems/top-k-frequent-elements/
/* Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order. */
func TopKFrequentElements(nums []int, k int) []int {
	note := make(map[int]int) // map { 1: 3, 2: 2, ...}
	for _, num := range nums {
		note[num]++
	}

	frequency := make([][]int, len(nums)+1) // slice [[], [3], [2], [1]]
	for num, count := range note {
		if frequency[count] == nil {
			frequency[count] = make([]int, 0)
		}

		frequency[count] = append(frequency[count], num)
	}

	result := make([]int, 0)
	for i := len(frequency) - 1; i >= 0; i-- {
		result = append(result, frequency[i]...)

		if len(result) >= k {
			break
		}
	}

	return result[:k]
}

// Related Problem: 451. Sort Characters By Frequency: https://leetcode.com/problems/sort-characters-by-frequency/
func frequencySort(s string) string {
	return FrequencySort(s)
}

// Related Problem: 692. Top K Frequent Words: https://leetcode.com/problems/top-k-frequent-words/
func topKFrequentWords(words []string, k int) []string {
	return TopKFrequent(words, k)
}
