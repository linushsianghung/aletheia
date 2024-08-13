package SlidingWindow

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// Ref: https://leetcode.com/problems/longest-substring-without-repeating-characters/solutions/2694302/js-98-sliding-window-with-exlanation/
/* Given a string s, find the length of the longest substring without repeating characters. */
func lengthOfLongestSubstring(s string) int {
	left, length := 0, 0
	note := make(map[rune]int)
	chars := []rune(s)

	for right := 0; right < len(s); right++ {
		if i, ok := note[chars[right]]; ok {
			if left <= i {
				left = i + 1
			}
		}
		note[chars[right]] = right
		if length < right-left+1 {
			length = right - left + 1
		}
	}

	return length
}
