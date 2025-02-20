package Medium

// https://leetcode.com/problems/longest-palindromic-substring/
// Ref: [Longest palindrome substring](https://www.youtube.com/watch?v=DK5OKKbF6GI)
/*
Given a string s, return the longest palindromic substring in s.

Analysis:
- racecar
- xxaabbaayy

_ _ _ _ i _ _ _
0 1 2 3 4 5 6 7
current i = 4

if length = 5 => (s, i, i)
start = 2 = 4 - (5-1)/2
end = 6 = 4 + length/2

if length = 6 => (s, i, i+1)
start = 2 = 4 - (6-1)/2
end = 7 = 4 + 6/2

*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var result string
	for i := 0; i < len(s)-1; i++ {
		palindrome1 := outwardComparison(s, i, i)
		palindrome2 := outwardComparison(s, i, i+1)

		if len(palindrome1) > len(result) || len(palindrome2) > len(result) {
			if len(palindrome1) > len(palindrome2) {
				result = palindrome1
			} else {
				result = palindrome2
			}
		}
	}

	return result
}

func longestPalindromeExercise(s string) string {
	return ""
}

func outwardComparison(s string, left, right int) string {
	var result string

	for left >= 0 && right < len(s) {
		// Based on the constraints, s consist of only digits and English letters. So it can just use "s" for iteration directly
		if s[left] != s[right] {
			break
		}

		result = s[left : right+1]
		left--
		right++
	}

	return result
}

func outwardComparisonExercise(s string, left, right int) string {
	return ""
}

func outwardComparisonAlt(s string, left, right int) int {
	length, sRune := 0, []rune(s)

	for left >= 0 && right < len(s) {
		if sRune[left] != sRune[right] {
			return length
		}

		length = right - left + 1
		left--
		right++
	}

	return length
}
