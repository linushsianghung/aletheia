package HashTable

// https://leetcode.com/problems/first-unique-character-in-a-string/description/
/* Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1. */
func firstUniqChar(s string) int {
	note := make(map[rune]int)

	for _, r := range s {
		note[r]++
	}

	for i, r := range s {
		if note[r] == 1 {
			return i
		}
	}

	return -1
}
