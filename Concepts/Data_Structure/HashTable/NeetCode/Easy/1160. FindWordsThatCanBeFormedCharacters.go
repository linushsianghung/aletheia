package Easy

import "maps"

// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/description/
/*
You are given an array of strings words and a string chars.

A string is good if it can be formed by characters from chars (each character can only be used once for each word in words).

Return the sum of lengths of all good strings in words.
*/
func countCharacters(words []string, chars string) int {
	note := make(map[rune]int)
	for _, c := range chars {
		note[c]++
	}

	count := 0
	for _, word := range words {
		temp := maps.Clone(note)
		isGood := true

		for _, w := range word {
			if _, ok := temp[w]; !ok {
				isGood = false
				break
			}

			temp[w]--
			if temp[w] == 0 {
				delete(temp, w)
			}
		}

		if isGood {
			count += len(word)
		}
	}

	return count
}
