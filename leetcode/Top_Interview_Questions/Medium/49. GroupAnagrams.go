package Medium

import "sort"

// GroupAnagrams https://leetcode.com/problems/group-anagrams/
/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Analysis:
Time Complexity: O(N * KlogK)
1. Outer Loop (N): The loop for _, str := range strs runs N times.
2. Inner Operation (Sorting): Inside the loop, you sort a string of length K.
	- Sorting takes O(KlogK)
	- Map access and hashing the string takes $O(K)$.
	- The dominant term inside the loop is the sorting: O(KlogK).
3. Total Complexity: $$N \times O(KlogK) = O(N * KlogK)
If you assume the string length K is small and constant (e.g., always less than 100), you could argue it approaches O(N), but strictly speaking, it depends on both variables.
*/
// GroupAnagrams implements the O(N * K) approach using character counts.
// In Go, arrays (like [26]int) can be used as map keys so instead of sorting (KlogK), we count the frequency of each character (which costs K).
// "eat" -> [1, 0, 0, 0, 1, ... 1 ...] (1 'a', 1 'e', 1 't')
// "tea" -> [1, 0, 0, 0, 1, ... 1 ...] (Same array)
func GroupAnagrams(strs []string) [][]string {
	// Key is an array of 26 integers (for a-z). Arrays are comparable in Go and can be map keys.
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		groups[count] = append(groups[count], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func GroupAnagramsNaive(strs []string) [][]string {
	note := make(map[string][]string)

	for _, str := range strs {
		sortedStr := []rune(str)
		sort.Slice(sortedStr, func(i, j int) bool {
			return sortedStr[i] < sortedStr[j]
		})

		key := string(sortedStr)
		note[key] = append(note[key], str)
		//if anagram, ok := note[string(sortedStr)]; ok {
		//	note[string(sortedStr)] = append(anagram, str)
		//} else {
		//	note[string(sortedStr)] = []string{str}
		//}
	}

	result := make([][]string, 0, len(note))
	for _, anagram := range note {
		result = append(result, anagram)
	}

	return result
}

func groupAnagramsExercise(strs []string) [][]string {
	return nil
}
