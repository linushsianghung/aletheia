package HashTable

// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/description/
// Reference: https://chatgpt.com/share/69d4e4b3-e160-8320-abaf-ea70b4a989a1
/*
You are given an array of strings words. Each element of words consists of two lowercase English letters.

Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. Each element can be selected at most once.

Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.

A palindrome is a string that reads the same forward and backward.
*/
func longestPalindrome(words []string) int {
	note := make(map[string]int)
	for _, w := range words {
		note[w]++
	}

	length, center := 0, false
	for word, count := range note {
		rev := string([]byte{word[1], word[0]})

		if word == rev {
			// Case: "aa", "bb"
			pairs := count / 2
			length += pairs * 4

			if count%2 == 1 {
				center = true
			}
		} else if word < rev {
			// Avoid double counting
			pairs := min(count, note[rev])
			length += pairs * 4
		}
	}

	if center {
		length += 2
	}

	return length
}
