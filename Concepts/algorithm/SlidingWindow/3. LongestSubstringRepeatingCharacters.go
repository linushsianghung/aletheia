package SlidingWindow

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
/* Given a string s, find the length of the longest substring without repeating characters. */
func lengthOfLongestSubstring(s string) int {
	winStart, maxLen := 0, 0
	sRune := []rune(s)
	note := make(map[rune]int)

	for winEnd := 0; winEnd < len(sRune); winEnd++ {
		// Dynamic-Size Sliding Window (Double For Loop): skip to the next character of the repeating one
		if i, ok := note[sRune[winEnd]]; ok && winStart <= i {
			winStart = i + 1
		}

		currentLen := winEnd - winStart + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
		note[sRune[winEnd]] = winEnd
	}

	return maxLen
}

func lengthOfLongestSubstringExercise(s string) int {
	return 0
}
