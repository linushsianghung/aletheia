package TwoPointers

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/
/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	r := []rune(s)

	for left < right {
		// In case of there are conti... non-digit and non-letter characters, it has to use if-else if to loop through all of them.
		if !unicode.IsLetter(r[left]) && !unicode.IsDigit(r[left]) {
			left++
		} else if !unicode.IsLetter(r[right]) && !unicode.IsDigit(r[right]) {
			right--
		} else {
			if r[left] != r[right] {
				return false
			}
			left++
			right--
		}
	}

	return true
}

func isPalindromeExercise(s string) bool {
	if len(s) == 0 {
		return true
	}

	return false
}
