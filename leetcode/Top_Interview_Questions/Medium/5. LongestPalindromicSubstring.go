package Medium

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions"

// https://leetcode.com/problems/longest-palindromic-substring/
// Ref: [Longest palindrome substring](https://www.youtube.com/watch?v=DK5OKKbF6GI)
/*
Given a string s, return the longest palindromic substring in s.

Analysis:
- racecar
- xxaabbaayy
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := divergeCheck(s, i, i)
		len2 := divergeCheck(s, i, i+1)

		length := Top_Interview_Questions.Max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	sRune := []rune(s)
	return string(sRune[start : end+1])
}

func divergeCheck(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		sRune := []rune(s)

		if sRune[left] == sRune[right] {
			left--
			right++
		} else {
			break
		}
	}

	return right - left + 1
}
