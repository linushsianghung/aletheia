package Sliding_Window

// LengthOfLongestSubstring https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
/*
Given a string s, find the length of the longest substring without repeating characters.

Analysis:
Time Complexity: O(n)
1. The range keyword iterates over the string s. Even though it decodes UTF-8 characters (runes) on the fly, it visits every character exactly once. If the string has n characters, this loop runs n times.
2. Inside the Loop (Constant Time Operations): Inside the loop, you perform the following operations:
	- Map Lookup: if lastIdx, ok := note[r]; — In Go, map lookups are $O(1)$ on average.
	- Comparison & Assignment: winStart <= lastIdx and winStart = ... — These are basic arithmetic operations, $O(1)$.
	- Math: max(...) — This is a simple comparison, $O(1)$.
	- Map Update: note[r] = winEnd — Map insertions are $O(1)$ on average.
3. Total Calculation: Since you perform a constant amount of work ($O(1)$) for each of the $n$ characters, the total time complexity is: $$n \times O(1) = O(n)$$

Space Complexity: O(min(n, A)), where $A$ is the size of the alphabet (e.g., 128 for ASCII, or more for Unicode).
	- The note map stores unique characters.
	- In the worst case (a string with all unique characters), the map grows to size $n$.
	- However, for a fixed alphabet (like English letters), the map size is capped (it won't grow infinitely), so it can also be considered $O(1)$ in limited character set scenarios.
*/
func LengthOfLongestSubstring(s string) int {
	winStart, maxLen := 0, 0
	// Window State: appeared characters and its position
	// Renaming note to lastSeen and anchor to lastIdx to describe what the data represents for instantly understandable to other engineers.
	lastSeen := make(map[rune]int)

	//for winEnd, r := range sRune {
	for winEnd, r := range s {
		// Dynamic-Size Sliding_Window: based on the index of repeating character
		// winStart <= anchor is necessary because it doesn't matter to the maxLen if the repeated character is located before the character of winStart, like cases s = "abba"
		if lastIdx, ok := lastSeen[r]; ok && winStart <= lastIdx {
			winStart = lastIdx + 1
		}

		maxLen = max(maxLen, winEnd-winStart+1)
		lastSeen[r] = winEnd
	}

	return maxLen
}

func lengthOfLongestSubstringExercise(s string) int {
	return 0
}
