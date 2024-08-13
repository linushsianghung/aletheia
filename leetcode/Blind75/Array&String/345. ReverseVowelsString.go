package Array_String

import "slices"

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envId=leetcode-75
/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/
func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	chars := []rune(s)
	vowels := []rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

	for left < right {
		if !slices.Contains(vowels, chars[left]) {
			left++
			continue
		}
		if !slices.Contains(vowels, chars[right]) {
			right--
			continue
		}

		temp := chars[left]
		chars[left] = chars[right]
		chars[right] = temp

		left++
		right--
	}

	return string(chars)
}
