package Easy

import "strings"

// https://leetcode.com/problems/word-pattern/description/
/*
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:

Each letter in pattern maps to exactly one unique word in s.
Each unique word in s maps to exactly one letter in pattern.
No two letters map to the same word, and no two words map to the same letter.
*/
func wordPattern(pattern string, s string) bool {
	noteP := make(map[byte]string)
	noteS := make(map[string]byte)

	strs := strings.Fields(s)
	if len(pattern) != len(strs) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if str, ok := noteP[pattern[i]]; ok && str != strs[i] {
			return false
		}
		noteP[pattern[i]] = strs[i]

		if p, ok := noteS[strs[i]]; ok && p != pattern[i] {
			return false
		}
		noteS[strs[i]] = pattern[i]
	}

	return true
}
