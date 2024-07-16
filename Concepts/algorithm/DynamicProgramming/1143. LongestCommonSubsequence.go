package DynamicProgramming

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

Example:
Input: text1 = "abcde", text2 = "ace"

## Identify sub-problems:
func lcs("abcde", "ace")
=> 1 + lcs("abcd", "ac")
=> 1 + max(
	lcs("abc", "ac") => max(
			lcs("ab", "ac") => ......
			lcs("abc", "a") => ......
	)
	lcs("abcd", "a") => max(
			lcs("abc", "a") => ......
			lcs("abcd", "") => return 0
	)
)

## DP Table:
Top-Down:
   a  b  c  d  e
a  3  2  2  1  1  0
c  2  2  2  1  1  0
e  1  1  1  1  1  0
   0  0  0  0  0  0

Bottom-Up:
   0  a  b  c  d  e
0  0  0  0  0  0  0
a  0  1  1  1  1  1
c  0  1  1  2  2  2
e  0  1  1  2  2  3

*/
func longestCommonSubsequence(text1 string, text2 string) int {
	return 0
}
