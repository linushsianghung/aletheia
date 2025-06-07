package SlidingWindow

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
/* Given a string s, find the length of the longest substring without repeating characters. */
func lengthOfLongestSubstring(s string) int {
	winStart, maxLen := 0, 0
	note := make(map[rune]int)

	//for winEnd, r := range sRune {
	for winEnd, r := range s {
		// Dynamic-Size Sliding Window: skip to the next character of the repeating one
		if i, ok := note[r]; ok && winStart <= i {
			winStart = i + 1
		}

		currentLen := winEnd - winStart + 1
		maxLen = max(currentLen, maxLen)
		note[r] = winEnd
	}

	return maxLen
}

func lengthOfLongestSubstringExercise(s string) int {
	return 0
}
