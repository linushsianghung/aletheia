package Medium

import "strings"

// https://leetcode.com/problems/word-break
/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

/*
	Complexity Analysis:
 1. Decision Tree / Recursion:
    At every step, we try to prefix a word from the dictionary. If it matches, we recursively check the remaining suffix.
    This leads to a tree with depth M (length of string).

2. Memorisation (Top-Down):
  - Time Complexity: O(m^2 * n) or O(m * n * k) where m = len(s), n = len(wordDict), k = max word length.
    We have 'm' states (starting positions). For each state, we iterate up to 'm' or 'k' to check words.
  - Space Complexity: O(m) for the memo map and recursion stack.

3. Tabulation (Bottom-Up):
  - Time Complexity: O(m^2 * n)
  - Space Complexity: O(m)
  - Concept: dp[i] represents if s[0:i] can be broken. To calculate dp[i], we look back at all dp[j] (where j < i). If dp[j] is true and s[j:i] is in the dictionary, then dp[i] is true.

4. How to optimize in an interview:
  - Use a Hash Set (map) for O(1) word lookups.
  - Store the length of the longest word in wordDict to limit the inner loop range (reducing O(m^2) to O(m * max_word_len)).
*/
func wordBreak(s string, wordDict []string) bool {
	return wordBreakTabulation(s, wordDict)
}

func wordBreakMemorisation(s string, wordDict []string) bool {
	// For efficient lookups, convert the word dictionary slice into a set (map).
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	// memo stores whether the suffix s[start:] can be broken
	memo := make(map[int]bool)

	var canBreak func(int) bool
	canBreak = func(start int) bool {
		// Base case: we reached the end of the string
		if start == len(s) {
			return true
		}
		if memo[start] {
			return memo[start]
		}

		// Try every possible end position for the first word
		for end := start + 1; end <= len(s); end++ {
			word := s[start:end]
			if wordSet[word] && canBreak(end) {
				memo[start] = true
				return true
			}
		}

		memo[start] = false
		return false
	}

	return canBreak(0)
}

func wordBreakMemorisationOptimised(s string, wordDict []string) bool {
	// For efficient lookups, convert the word dictionary slice into a set (map).
	wordSet := make(map[string]bool)
	maxLen := 0
	for _, word := range wordDict {
		wordSet[word] = true
		maxLen = max(maxLen, len(word))
	}

	// Use a slice instead of a map for memoization.
	// 0: unvisited, 1: true, 2: false
	memo := make([]int, len(s))

	var canBreak func(int) bool
	canBreak = func(start int) bool {
		// Base case: we reached the end of the string
		if start == len(s) {
			return true
		}
		if memo[start] != 0 {
			return memo[start] == 1
		}

		// Optimization: Only check up to maxLen or the end of the string
		for end := start + 1; end <= len(s) && end-start <= maxLen; end++ {
			word := s[start:end]
			if wordSet[word] && canBreak(end) {
				memo[start] = 1
				return true
			}
		}

		memo[start] = 2
		return false
	}

	return canBreak(0)
}

// It would fail with Time Limit Exceeded (TLE)
func wordBreakStringBase(s string, wordDict []string) bool {
	memo := make(map[string]bool)

	var canBreak func(target string) bool
	canBreak = func(target string) bool {
		if len(target) == 0 {
			return true
		}

		if memo[target] {
			return memo[target]
		}

		for _, word := range wordDict {
			if strings.HasPrefix(target, word) {
				if canBreak(target[len(word):]) {
					memo[target] = true
					return true
				}
			}
		}

		memo[target] = false
		return false
	}

	return canBreak(s)
}

func wordBreakTabulation(s string, wordDict []string) bool {
	// 1. For efficient lookups, convert the word dictionary slice into a set (map).
	wordSet := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, word := range wordDict {
		wordSet[word] = true
		maxLen = max(maxLen, len(word))
	}

	// 2. Create a DP table. dp[i] will be true if the prefix s[:i] can be segmented.
	// A common and cleaner convention for this type of problem is to use a DP table of size n+1. This allows dp[0] to represent the base case (an empty string), which is always true.
	dp := make([]bool, len(s)+1)
	dp[0] = true // Base case: An empty string can always be segmented.

	// 3. Iterate through the string to fill the DP table.
	// `i` represents the length of the prefix we are checking (s[:i]).
	for i := 1; i <= len(s); i++ {
		// `j` is the split point. We check if s[:j] is breakable and s[j:i] is a word.
		// Optimization: j only needs to go back as far as the longest word.
		start := 0
		if i > maxLen {
			start = i - maxLen
		}
		for j := i - 1; j >= start; j-- {
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

func wordBreakTabulationSimplified(s string, wordDict []string) bool {
	// DP Table: note[i] means s[0...i] is breakable
	note := make([]bool, len(s))

	for i := range len(s) {
		for _, word := range wordDict {
			// Checking if a substring ending at index i matches 'word'
			if len(word) <= i+1 && word == s[i-len(word)+1:i+1] {
				// Check if the prefix before this word was valid
				if i-len(word) < 0 || note[i-len(word)] {
					note[i] = true
					break
				}
			}
		}
	}

	return note[len(s)-1]
}

func wordBreakExercise(s string, wordDict []string) bool {
	return false
}

// Related Problem: 91. Decode Ways: https://leetcode.com/problems/decode-ways/description/
func numDecodings(s string) int {
	return NumDecodings(s)
}
