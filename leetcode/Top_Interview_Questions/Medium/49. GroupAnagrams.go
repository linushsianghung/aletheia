package Medium

import "sort"

// https://leetcode.com/problems/group-anagrams/
/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/
func groupAnagrams(strs []string) [][]string {
	note := make(map[string][]string)

	for _, str := range strs {
		sortedRune := []rune(str)
		sort.Slice(sortedRune, func(i, j int) bool {
			return sortedRune[i] < sortedRune[j]
		})

		if s, ok := note[string(sortedRune)]; ok {
			note[string(sortedRune)] = append(s, str)
		} else {
			note[string(sortedRune)] = []string{str}
		}
	}

	result := make([][]string, 0)
	for _, v := range note {
		result = append(result, v)
	}

	return result
}

func groupAnagramsExercise(strs []string) [][]string {
	return nil
}
