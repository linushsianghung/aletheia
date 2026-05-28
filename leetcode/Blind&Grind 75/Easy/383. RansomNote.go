package Easy

// https://leetcode.com/problems/ransom-note/description
/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.
*/
func canConstruct(ransomNote string, magazine string) bool {
	var ransomNoteArray [26]int

	for _, r := range ransomNote {
		ransomNoteArray[r-'a']++
	}

	for _, r := range magazine {
		ransomNoteArray[r-'a']--
	}

	for _, count := range ransomNoteArray {
		if count > 0 {
			return false
		}
	}

	return true
}

func canConstructExercise(ransomNote string, magazine string) bool {

	return true
}
