package SlidingWindow

import "slices"

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envId=leetcode-75
/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
*/
func maxVowels(s string, k int) int {
	windowStart, count, maxCount := 0, 0, 0
	sRune := []rune(s)
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

	for windowEnd := range s {
		if slices.Contains(vowels, sRune[windowEnd]) {
			count++
		}
		if windowEnd-windowStart == k-1 {
			maxCount = max(maxCount, count)

			if slices.Contains(vowels, sRune[windowStart]) {
				count--
			}
			windowStart++
		}
	}

	return maxCount
}
