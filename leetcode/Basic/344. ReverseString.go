package Basic

import "fmt"

// https://leetcode.com/problems/reverse-string/description/
/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.
*/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseStringExercise(s []byte) {

}

// Unsolved...
func reverseStringRecursively(s []byte) []byte {
	var localReverseFunc func(s []byte) []byte
	localReverseFunc = func(s []byte) []byte {
		if len(s) == 0 {
			return []byte{}
		}
		return append(localReverseFunc(s[1:]), s[0])
	}

	result := localReverseFunc(s)
	fmt.Println(s)
	fmt.Println(result)
	return result
}
