package Medium

// https://leetcode.com/problems/word-break
/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/
func wordBreak(s string, wordDict []string) bool {
	// DP Table
	note := make([]bool, len(s))

	for i := range len(s) {
		for _, word := range wordDict {
			// Check if word is equal to substring
			if len(word) <= i+1 && word == s[i-len(word)+1:i+1] {
				// Check if last part of substring is valid
				if i-len(word) < 0 || note[i-len(word)] {
					note[i] = true
					break
				}
			}
		}
	}

	return note[len(s)-1]
}
