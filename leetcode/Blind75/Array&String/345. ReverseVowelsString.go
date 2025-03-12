package Array_String

import (
	"slices"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envId=leetcode-75
/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/
func reverseVowels(s string) string {
	sRune := []rune(s)
	vowel := []rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

	left, right := 0, len(sRune)-1
	for left < right {
		if !slices.Contains(vowel, sRune[left]) {
			left++
			continue
		}
		if !slices.Contains(vowel, sRune[right]) {
			right--
			continue
		}

		sRune[left], sRune[right] = sRune[right], sRune[left]
		left++
		right--
	}

	return string(sRune)
}

func reverseVowelsExercise(s string) string {
	return ""
}
