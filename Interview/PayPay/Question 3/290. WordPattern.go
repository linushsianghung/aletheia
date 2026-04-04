package Question_3

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
	notePS := make(map[byte]string)
	noteSP := make(map[string]byte)

	strs := strings.Fields(s)
	if len(strs) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if str, ok := notePS[pattern[i]]; ok && str != strs[i] {
			return false
		}
		notePS[pattern[i]] = strs[i]

		if p, ok := noteSP[strs[i]]; ok && p != pattern[i] {
			return false
		}
		noteSP[strs[i]] = pattern[i]
	}

	return true
}
