package Array_String

// https://leetcode.com/problems/merge-strings-alternately/
/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.
*/
func mergeAlternately(word1 string, word2 string) string {
	runner1, runner2 := 0, 0
	r1 := []rune(word1)
	r2 := []rune(word2)
	word := make([]rune, 0)

	for runner1 < len(word1) && runner2 < len(word2) {
		word = append(word, r1[runner1])
		word = append(word, r2[runner2])
		runner1++
		runner2++
	}

	if runner1 < len(word1) {
		word = append(word, r1[runner1:]...)
	}

	if runner2 < len(word2) {
		word = append(word, r2[runner2:]...)
	}

	return string(word)
}
