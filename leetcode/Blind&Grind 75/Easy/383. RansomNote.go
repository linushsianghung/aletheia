package Easy

// https://leetcode.com/problems/ransom-note/description
/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.
*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var ransomArray, magazineArray [26]int

	for _, c := range ransomNote {
		ransomArray[c-'a']++
	}
	for _, c := range magazine {
		magazineArray[c-'a']++
	}

	for i, count := range ransomArray {
		if magazineArray[i] < count {
			return false
		}
	}

	return true
}
