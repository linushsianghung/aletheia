package Medium

// https://leetcode.com/problems/palindromic-substrings/
/*
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.
*/
func countSubstrings(s string) int {
	count := 0

	for i := range s {
		count += countPalindrome(s, i, i)
		count += countPalindrome(s, i, i+1)
	}

	return count
}

func countSubstringsExercise(s string) int {
	return 0
}

func countPalindrome(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}

		count++
		left--
		right++
	}

	return count
}

func countPalindromeExercise(s string, left, right int) int {
	return 0
}
