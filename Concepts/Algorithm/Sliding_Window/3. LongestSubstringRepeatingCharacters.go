package Sliding_Window

// LengthOfLongestSubstring https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
/* Given a string s, find the length of the longest substring without repeating characters. */
func LengthOfLongestSubstring(s string) int {
	winStart, maxLen := 0, 0
	// Window State: appeared characters and its position
	note := make(map[rune]int)

	//for winEnd, r := range sRune {
	for winEnd, r := range s {
		// Dynamic-Size Sliding_Window: based on the index of repeating character
		// winStart <= anchor is necessary because it doesn't matter to the maxLen if the repeated character is located before the character of winStart, like cases s = "abba"
		if anchor, ok := note[r]; ok && winStart <= anchor {
			winStart = anchor + 1
		}

		maxLen = max(maxLen, winEnd-winStart+1)
		note[r] = winEnd
	}

	return maxLen
}

func lengthOfLongestSubstringExercise(s string) int {
	return 0
}
