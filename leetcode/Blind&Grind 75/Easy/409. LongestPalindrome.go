package Easy

// https://leetcode.com/problems/longest-palindrome/description
/*
Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome.
*/
func longestPalindrome(s string) int {
	note := make(map[rune]int)

	for _, r := range s {
		note[r]++
	}

	length := 0
	hasOdd := false
	for _, count := range note {
		if count%2 == 0 {
			length += count
		} else {
			length += count - 1
			hasOdd = true
		}
	}

	if hasOdd {
		length += 1
	}

	return length
}

func longestPalindromeExercise(s string) int {
	return 0
}
