package Basic

import "math"

// Reverse https://leetcode.com/problems/reverse-integer/
/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
func Reverse(x int) int {
	rev := 0
	// It has to be non-equal because of negative cases, like -123 to -321
	for x != 0 {
		rev = rev*10 + x%10
		if rev > math.MaxInt32 || rev < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return rev
}

func reverseExercise(x int) int {
	return 0
}
