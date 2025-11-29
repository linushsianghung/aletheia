package Medium

import "sort"

// GroupAnagrams https://leetcode.com/problems/group-anagrams/
/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/
func GroupAnagrams(strs []string) [][]string {
	note := make(map[string][]string)

	for _, str := range strs {
		sortedStr := []rune(str)
		sort.Slice(sortedStr, func(i, j int) bool {
			return sortedStr[i] < sortedStr[j]
		})

		if anagram, ok := note[string(sortedStr)]; ok {
			note[string(sortedStr)] = append(anagram, str)
		} else {
			note[string(sortedStr)] = []string{str}
		}
	}

	result := make([][]string, 0)
	for _, anagram := range note {
		result = append(result, anagram)
	}

	return result
}

func groupAnagramsExercise(strs []string) [][]string {
	return nil
}
