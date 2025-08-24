package HashMap_Set

import (
	"sort"
)

// https://leetcode.com/problems/determine-if-two-strings-are-close/description/?envId=leetcode-75
// Reference:
// - https://leetcode.com/problems/determine-if-two-strings-are-close/solutions/2868036/python-c-o-n-one-liner-proof-explained
// - https://leetcode.com/problems/determine-if-two-strings-are-close/solutions/4561554/beats-100-c-java-python-js-explained-with-video-hash-sort-count/
/*
Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
- For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
- For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

Analysis:
The idea behind this solution consists in considering its invariants (i.e., properties that do not change after either of 2 operations was applied):
- New (unique) characters can not appear, neither can old ones be completely eliminated.
- Character frequencies can not be altered, they can only be arbitrarily redistributed (permuted) between existing characters.

This leads to a simple solution that consists in
- Checking that unique characters in both strings coincide.
- Checking that collections of character frequencies in both strings coincide.

*/
func closeStrings(word1 string, word2 string) bool {
	frequency1 := make([]int, 26)
	frequency2 := make([]int, 26)

	for _, c := range word1 {
		frequency1[c-'a']++
	}
	for _, c := range word2 {
		frequency2[c-'a']++
	}
	for i := range 26 {
		if (frequency1[i] == 0 && frequency2[i] != 0) || (frequency1[i] != 0 && frequency2[i] == 0) {
			return false
		}
	}

	sort.Ints(frequency1)
	sort.Ints(frequency2)
	for i := range 26 {
		if frequency1[i] != frequency2[i] {
			return false
		}
	}

	return true
}

func closeStringsExercise(word1 string, word2 string) bool {
	return false
}
