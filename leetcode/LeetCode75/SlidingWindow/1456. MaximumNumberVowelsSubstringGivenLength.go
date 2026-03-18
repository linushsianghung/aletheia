package SlidingWindow

import "github.com/linushung/aletheia/Concepts/Algorithm/Sliding_Window"

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envId=leetcode-75
/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
*/
func maxVowels(s string, k int) int {
	return Sliding_Window.MaxVowels(s, k)
}
