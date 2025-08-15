package Dynamic_Programming

import "fmt"

// LongestCommonSubsequence https://leetcode.com/problems/longest-common-subsequence/description/
// Ref:
// - Back to Back SWE: https://www.youtube.com/watch?v=ASoaQq66foQ
// - NeetCode: https://neetcode.io/solutions/longest-common-subsequence
/*
Given two strings textext1 and textext2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
For example, "ace" is a subsequence of "abcde".

A common subsequence of two strings is a subsequence that is common to both strings.

## Identify sub-problems:
func lcs("abcde", "ace")
=> 1 + lcs("abcd", "ac")
=> 1 + max(
	lcs("abc", "ac") => max(
			lcs("ab", "ac") => ......
			lcs("abc", "a") => ......
	)
	lcs("abcd", "a") => max(
			lcs("abc", "a") => return 1 + max(......)
			lcs("abcd", "") => return 0
	)
)

## Transfer to DP Table
- When using DP Table, always add 1 more row and column
- Set default value 0 (in this case) for the 1st row and column

   *  a  b  c  d  e
*  0  0  0  0  0  0
a  0  1  1  1  1  1
c  0  1  1  2  2  2
e  0  1  1  2  2  3

*/
func LongestCommonSubsequence(text1 string, text2 string) int {
	return lcsDPTabulationHelper(text1, text2)
	//return lcsDPMemorisationHelper(text1, text1, make(map[string]int))
}

/*
Time Complexity: O(mn)
Space Complexity: O(mn)
where m & n is the length of the 2 strings
*/
func lcsDPTabulationHelper(text1, text2 string) int {
	table := make([][]int, len(text1)+1)
	for i := range table {
		table[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[0]); j++ {
			if text1[i-1] == text2[j-1] {
				table[i][j] = 1 + table[i-1][j-1]
			} else {
				table[i][j] = max(table[i][j-1], table[i-1][j])
			}
		}
	}

	return table[len(text1)][len(text2)]
}

// "fatal error: runtime: cannot allocate memory runtime stack" will occur in quite longer string
// It might be because the note requires too many memory space
func lcsDPMemorisationHelper(text1, text2 string, note map[string]int) int {
	if count, ok := note[fmt.Sprintf("%s-%s", text1, text2)]; ok {
		return count
	}

	len1, len2 := len(text1), len(text2)

	if len1 == 0 || len2 == 0 {
		return 0
	}
	if text1[len1-1] == text2[len2-1] {
		key := fmt.Sprintf("%s-%s", text1, text2)
		note[key] = lcsDPMemorisationHelper(text1[:len1-1], text2[:len2-1], note)
		return 1 + note[key]
	}

	key1 := fmt.Sprintf("%s-%s", text1[:len1-1], text2)
	note[key1] = lcsDPMemorisationHelper(text1[:len1-1], text2, note)
	key2 := fmt.Sprintf("%s-%s", text1, text2[:len2-1])
	note[key2] = lcsDPMemorisationHelper(text1, text2[:len2-1], note)

	return max(note[key1], note[key2])
}
