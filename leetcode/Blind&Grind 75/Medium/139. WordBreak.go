package Medium

// https://leetcode.com/problems/word-break
/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/
func wordBreak(s string, wordDict []string) bool {
	return wordBreakDP(s, wordDict)
}

/*
The main idea is: dp[i] is true if s[:i] can be broken.
To calculate dp[i], we check for a split point j where dp[j] is true AND the substring s[j:i] is a valid word.
*/
func wordBreakDP(s string, wordDict []string) bool {
	// 1. For efficient lookups, convert the word dictionary slice into a set (map).
	// This changes word lookups from O(L) (linear scan) to O(1) on average.
	wordSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// 2. Create a DP table. dp[i] will be true if the prefix s[:i] can be segmented.
	// A common and cleaner convention for this type of problem is to use a DP table of size n+1. This allows dp[0] to represent the base case (an empty string), which is always true.
	dp := make([]bool, len(s)+1)
	dp[0] = true // Base case: An empty string can always be segmented.

	// 3. Iterate through the string to fill the DP table.
	// `i` represents the length of the prefix we are checking (s[:i]).
	for i := 1; i <= len(s); i++ {
		// `j` is the split point. We check if s[:j] is breakable and s[j:i] is a word.
		for j := 0; j < i; j++ {
			// If dp[j] is true, it means the prefix s[:j] is valid.
			// Now, check if the remaining part s[j:i] is a word in our set.
			if dp[j] && wordSet[s[j:i]] {
				// If both conditions are met, we've found a way to segment s[:i].
				dp[i] = true
				break // Move to the next `i` since we've confirmed s[:i] is possible.
			}
		}
	}

	// The result is whether the entire string (length `n`) can be segmented.
	return dp[len(s)]
}

/*
The biggest bottleneck is iterating through the entire wordDict for every single character in s.
If s has length N and wordDict has M words, this leads to a high time complexity.
*/
func wordBreakAlternative(s string, wordDict []string) bool {
	// DP Table
	note := make([]bool, len(s))

	for i := range len(s) {
		for _, word := range wordDict {
			// Checking if a substring ending at index i is a valid word
			if len(word) <= i+1 && word == s[i-len(word)+1:i+1] {
				// And then checking if the part before it was also valid (note[i-len(word)]),
				if i-len(word) < 0 || note[i-len(word)] {
					note[i] = true
					break
				}
			}
		}
	}

	return note[len(s)-1]
}

// Related Problem: 91. Decode Ways: https://leetcode.com/problems/decode-ways/description/
func numDecodings(s string) int {
	return numDecodingsMemo(s)
}
