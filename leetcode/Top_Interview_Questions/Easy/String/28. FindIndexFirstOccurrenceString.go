package String

import "strings"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
/* Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack. */
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	// It has to be cached to prevent the length shrinking during the operation
	haystackLen := len(haystack)
	for i := 0; i <= haystackLen-len(needle); i++ {
		if strings.HasPrefix(haystack[i:], needle) {
			return i
		}
	}

	return -1
}

func strStrExercise(haystack string, needle string) int {
	return 0
}
