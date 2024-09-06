package String

import "strings"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
/* Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack. */
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	haystackLen := len(haystack)
	for i := 0; i <= haystackLen-len(needle); i++ {
		if strings.HasPrefix(haystack, needle) {
			return i
		}

		haystack = haystack[1:]
	}

	return -1
}

func strStrExercise(haystack string, needle string) int {

	return 0
}
