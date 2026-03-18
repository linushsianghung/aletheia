package Easy

// https://leetcode.com/problems/longest-palindrome/description
/*
Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome.
*/
func longestPalindrome(s string) int {
	note := make(map[rune]int)
	oddCount := 0

	for _, c := range s {
		note[c]++
		if note[c]%2 != 0 {
			oddCount++
		} else {
			oddCount--
		}
	}

	if oddCount > 0 {
		return len(s) - oddCount + 1
	} else {
		return len(s)
	}
}
