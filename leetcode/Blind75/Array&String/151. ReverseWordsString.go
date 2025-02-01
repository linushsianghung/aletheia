package Array_String

import (
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string/description/envId=leetcode-75
// Ref: https://leetcode.com/problems/reverse-words-in-a-string/solutions/47720/clean-java-two-pointers-solution-no-trim-no-split-no-stringbuilder/?envType=study-plan-v2&envId=leetcode-75
/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words. Do not include any extra spaces.
*/
func reverseWords(s string) string {
	// Using Fields function to break down the string into smaller substrings wherever one or more whitespace characters appear.
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		//strings.Trim(words[i], " ")
		//strings.Trim(words[j], " ")
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
