package DP_Multidimensional

import "fmt"

// https://leetcode.com/problems/edit-distance/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/edit-distance/solutions/159295/python-solutions-and-intuition
/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character

Analysis:

base case: word1 = "" or word2 = "" => return length of other string
recursive case: word1[0] == word2[0] => recurse on word1[1:] and word2[1:]
recursive case: word1[0] != word2[0] => recurse by inserting, deleting, or replacing
*/
func minDistance(word1 string, word2 string) int {
	return minDistanceHelper(word1, word2, 0, 0, make(map[string]int))
}

func minDistanceHelper(word1, word2 string, index1, index2 int, memo map[string]int) int {
	if len(word1[index1:]) == 0 {
		return len(word2[index2:])
	}
	if len(word2[index2:]) == 0 {
		return len(word1[index1:])
	}

	if count, ok := memo[fmt.Sprintf("%s-%s", word1[index1:], word2[index2:])]; ok {
		return count
	}

	result := 0
	if word1[index1] == word2[index2] {
		result = minDistanceHelper(word1, word2, index1+1, index2+1, memo)
	} else {
		insert := 1 + minDistanceHelper(word1, word2, index1, index2+1, memo)
		remove := 1 + minDistanceHelper(word1, word2, index1+1, index2, memo)
		replace := 1 + minDistanceHelper(word1, word2, index1+1, index2+1, memo)
		result = min(insert, min(remove, replace))
	}

	memo[fmt.Sprintf("%s-%s", word1[index1:], word2[index2:])] = result
	return result
}
