package HashTable

// FrequencySort https://leetcode.com/problems/sort-characters-by-frequency/description/
/*
Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.

Return the sorted string. If there are multiple answers, return any of them.
*/
func FrequencySort(s string) string {
	note := make(map[rune]int) //  map { e: 2, r: 1, ...}
	for _, c := range s {
		note[c]++
	}

	frequency := make([][]rune, len(s)+1) // [[], [r, t], [e,e,...], ...]
	for c, count := range note {
		if frequency[count] == nil {
			frequency[count] = make([]rune, 0)
		}

		for range count {
			frequency[count] = append(frequency[count], c)
		}
	}

	result := make([]rune, 0)
	for i := len(frequency) - 1; i >= 0; i-- {
		result = append(result, frequency[i]...)
	}

	return string(result)
}
