package TwoPointers

import (
	"strings"
	"unicode"
)

// IsPalindrome https://leetcode.com/problems/valid-palindrome/
/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	sRune := []rune(s)

	for left < right {
		if !unicode.IsLetter(sRune[left]) && !unicode.IsDigit(sRune[left]) {
			left++
			continue
		}
		if !unicode.IsLetter(sRune[right]) && !unicode.IsDigit(sRune[right]) {
			right--
			continue
		}

		if sRune[left] != sRune[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isPalindromeExercise(s string) bool {
	return false
}
