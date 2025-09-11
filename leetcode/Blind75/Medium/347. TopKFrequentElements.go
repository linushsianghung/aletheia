package Medium

// https://leetcode.com/problems/top-k-frequent-elements/
/* Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order. */
func topKFrequent(nums []int, k int) []int {
	note := make(map[int]int)
	for _, num := range nums {
		note[num]++
	}

	frequency := make([][]int, len(nums)+1)
	for key, value := range note {
		if frequency[value] == nil {
			frequency[value] = make([]int, 0)
		}

		frequency[value] = append(frequency[value], key)
	}

	count, result := 0, make([]int, 0)
	for i := len(frequency) - 1; i >= 0; i-- {
		if count >= k {
			break
		}

		result = append(result, frequency[i]...)
		count += len(frequency[i])
	}

	return result
}

// Related Problem: 451. Sort Characters By Frequency: https://leetcode.com/problems/sort-characters-by-frequency/
func frequencySort(s string) string {
	note := make(map[rune]int)
	for _, r := range s {
		note[r]++
	}

	frequency := make([][]rune, len(s)+1)
	for key, value := range note {
		if frequency[value] == nil {
			frequency[value] = make([]rune, 0)
		}

		for range value {
			frequency[value] = append(frequency[value], key)
		}
	}

	result := make([]rune, 0)
	for i := len(frequency) - 1; i >= 0; i-- {
		result = append(result, frequency[i]...)
	}

	return string(result)
}
