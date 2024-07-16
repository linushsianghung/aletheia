package Basic

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/palindrome-number/
/* Given an integer x, return true if x is a palindrome, and false otherwise. */
func isPalindrome(x int) bool {
	/* Basic check */
	if x == 0 {
		return true
	}

	/* Negative cases check */
	if x < 0 || x%10 == 0 {
		return false
	}

	return isPalindromeImpl1(x)
	// return isPalindromeImpl2(x)
}

func isPalindromeImpl1(x int) bool {
	/* Half-way reverse: the x will larger than reverse(popped up digits) until the number of digits less than reverse  */
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}

	if rev == x || rev/10 == x {
		return true
	}

	return false
}

func isPalindromeImpl1Exercise(x int) bool {
	return false
}

func isPalindromeImpl2(x int) bool {
	xSli := strings.Split(strconv.Itoa(x), "")
	i, j := 0, len(xSli)-1
	for i <= j {
		if xSli[i] != xSli[j] {
			return false
		}
		i++
		j--
	}

	return true
}
