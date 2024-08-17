package TwoPointers

// https://leetcode.com/problems/is-subsequence/description/?envId=leetcode-75
/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters
without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	sIdx := 0
	sRune, tRune := []rune(s), []rune(t)
	for _, r := range tRune {
		if r == sRune[sIdx] {
			sIdx++
			if sIdx == len(sRune) {
				return true
			}
		}
	}

	return sIdx == len(sRune)
}
