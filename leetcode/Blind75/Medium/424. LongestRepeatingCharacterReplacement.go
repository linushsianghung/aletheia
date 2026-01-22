package Medium

// https://leetcode.com/problems/longest-repeating-character-replacement/description/
/*
You are given a string s and an integer k.
You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.
*/
func characterReplacement(s string, k int) int {
	winStart, maxCount, maxLen := 0, 0, 0
	note := make(map[rune]int)
	sRune := []rune(s)

	for winEnd, r := range sRune {
		note[r]++
		maxCount = max(maxCount, note[r])

		if winEnd-winStart+1-maxCount > k {
			note[sRune[winStart]]--
			winStart++
		}
		maxLen = max(maxLen, winEnd-winStart+1)
	}

	return maxLen
}
