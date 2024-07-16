package String

import (
	"strings"
)

// https://leetcode.com/problems/longest-common-prefix/
// Ref: https://www.youtube.com/watch?v=H8A9twm07Vc
/* Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string "". */
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	return longestCommonPrefixHorizon(strs)
}

// Horizontal Scanning which is quicker for big number of words
func longestCommonPrefixHorizon(strs []string) string {

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]

			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

func longestCommonPrefixHorizonExercise(strs []string) string {
	return ""
}

// Vertical Scanning which is quicker for limited number of words
func longestCommonPrefixVertical(strs []string) string {
	return ""
}
