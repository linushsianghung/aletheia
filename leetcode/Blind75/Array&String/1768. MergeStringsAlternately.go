package Array_String

// https://leetcode.com/problems/merge-strings-alternately/
/*
You are given two strings word1 and word2. Merging the strings by adding letters in alternating order, starting with word1.
If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.
*/
func mergeAlternately(word1 string, word2 string) string {
	result := make([]rune, 0)
	runner1, runner2 := 0, 0
	rune1, rune2 := []rune(word1), []rune(word2)

	for runner1 < len(word1) && runner2 < len(word2) {
		result = append(result, rune1[runner1])
		result = append(result, rune2[runner2])

		runner1++
		runner2++
	}

	if runner1 < len(word1) {
		result = append(result, rune1[runner1:]...)
	}
	if runner2 < len(word2) {
		result = append(result, rune2[runner2:]...)
	}

	return string(result)
}

func mergeAlternatelyExercise(word1 string, word2 string) string {
	return ""
}
