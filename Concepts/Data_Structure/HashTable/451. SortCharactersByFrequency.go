package HashTable

// FrequencySort https://leetcode.com/problems/sort-characters-by-frequency/description/
/*
Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.

Return the sorted string. If there are multiple answers, return any of them.
*/
func FrequencySort(s string) string {
	note := make(map[rune]int)
	for _, r := range s {
		note[r]++
	}

	frequency := make([][]rune, len(s)+1)
	for key, value := range note {
		if frequency[value] == nil {
			frequency[value] = make([]rune, 0)
		}

		// Using for loop to put the according amount of character
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
