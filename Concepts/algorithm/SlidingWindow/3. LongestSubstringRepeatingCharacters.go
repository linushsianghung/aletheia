package SlidingWindow

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// Ref: https://leetcode.com/problems/longest-substring-without-repeating-characters/solutions/2694302/js-98-sliding-window-with-exlanation/
/* Given a string s, find the length of the longest substring without repeating characters. */
func lengthOfLongestSubstring(s string) int {
	length := 0
	note := make(map[rune]int)
	chars := []rune(s)

	max := func(x, y int) int {
		if x >= y {
			return x
		} else {
			return y
		}
	}

	for slow, fast := 0, 0; fast < len(s); fast++ {
		if i, ok := note[chars[fast]]; ok {
			if slow <= i {
				slow = i + 1
			}
		}
		note[chars[fast]] = fast
		length = max(length, fast-slow+1)
	}

	return length
}
