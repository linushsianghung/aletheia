package Easy

// https://leetcode.com/problems/valid-anagram/
/* Given two strings s and t, return true if t is an anagram of s, and false otherwise. */
func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	tRune := []rune(t)
	note := make(map[rune]int)
	for i, r := range s {
		note[r]++
		note[tRune[i]]--
	}

	for _, val := range note {
		if val != 0 {
			return false
		}
	}

	return true
}

func isAnagramExercise(s string, t string) bool {
	return false
}

func isAnagramAlt(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	note := make(map[rune]int)
	for _, r := range s {
		note[r]++
	}

	for _, r := range t {
		if _, ok := note[r]; !ok {
			return false
		}
		note[r]--
		if note[r] == 0 {
			delete(note, r)
		}
	}

	return len(note) == 0
}
