package DynamicProgramming

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions"

// https://leetcode.com/problems/longest-common-subsequence/description/
// Ref:
// - Back to Back SWE: https://www.youtube.com/watch?v=ASoaQq66foQ
// - NeetCode: https://www.youtube.com/watch?v=Ua0GhsJSlWM
// - Abdul Bari: https://www.youtube.com/watch?v=sSno9rV8Rhg
/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

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

   0  a  b  c  d  e
0  0  0  0  0  0  0
a  0  1  1  1  1  1
c  0  1  1  2  2  2
e  0  1  1  2  2  3

*/
func longestCommonSubsequence(text1 string, text2 string) int {
	return lcsDPTableHelper(text1, text2)
}

/*
Time Complexity: O(mn)
Space Complexity: O(mn)
where m & n is the length of the 2 strings
*/
// Time Limit Exceeded in quite longer strings
func lcsDPTableHelper(t1, t2 string) int {
	table := make([][]int, len(t1)+1)
	for i := range table {
		table[i] = make([]int, len(t2)+1)
	}

	for i := 0; i < len(t1); i++ {
		for j := 0; j < len(t2); j++ {
			r1, r2 := []rune(t1), []rune(t2)
			if r1[i] == r2[j] {
				table[i+1][j+1] = 1 + table[i][j]
			} else {
				table[i+1][j+1] = Top_Interview_Questions.Max(table[i][j+1], table[i+1][j])
			}
		}
	}

	return table[len(t1)-1][len(t2)-1]
}

// Time Limit Exceeded in longer strings
func lcsRecursiveHelper(t1, t2 string) int {
	len1, len2 := len(t1), len(t2)

	if len1 == 0 || len2 == 0 {
		return 0
	}
	if t1[len1-1] == t2[len2-1] {
		return 1 + lcsRecursiveHelper(t1[:len1-1], t2[:len2-1])
	}

	return Top_Interview_Questions.Max(lcsRecursiveHelper(t1[:len1-1], t2), lcsRecursiveHelper(t1, t2[:len2-1]))
}
