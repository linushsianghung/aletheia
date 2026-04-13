package Easy

import "math"

// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/description/
/*
You are given a string s consisting of lowercase English letters.

Your task is to find the maximum difference diff = freq(a1) - freq(a2) between the frequency of characters a1 and a2 in the string such that:

a1 has an odd frequency in the string.
a2 has an even frequency in the string.
Return this maximum difference.
*/
func maxDifference(s string) int {
	frequency := make([]int, 26)
	for _, c := range s {
		frequency[c-'a']++
	}

	maxOdd, minEven := 0, math.MaxInt
	for _, count := range frequency {
		if count%2 == 0 {
			minEven = min(minEven, count)
		}
		if count%2 != 0 {
			maxOdd = max(maxOdd, count)
		}
	}

	return maxOdd - minEven
}
