package Easy

// https://leetcode.com/problems/isomorphic-strings/description
// Reference: https://chatgpt.com/share/69dc3ebe-3a5c-8322-b3ea-1b9d79a29410
/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	frequencyS, frequencyT := make([]int, 128), make([]int, 128)
	for i := 0; i < len(s); i++ {
		if frequencyS[s[i]] != frequencyT[t[i]] {
			return false
		}

		frequencyS[s[i]] = i + 1
		frequencyT[t[i]] = i + 1
	}

	return true
}
