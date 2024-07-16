package testdome

import (
	"fmt"
	"strings"
)

/**
Write a function that checks if a given word is a palindrome. Character case should be ignored.
*/
func isPalindrome(s string) bool {
	pali := strings.Split(strings.ToLower(s), "")

	low := 0
	high := len(pali) - 1

	for low <= high {
		if pali[low] != pali[high] {
			return false
		}
		low++
		high--
	}

	return true
}

func Palindrome() {
	fmt.Println(isPalindrome("Deleveled"))
}
