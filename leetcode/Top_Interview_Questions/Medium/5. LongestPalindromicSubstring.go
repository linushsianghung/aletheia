package Medium

// LongestPalindrome https://leetcode.com/problems/longest-palindromic-substring/
// Ref: [Longest palindrome substring](https://www.youtube.com/watch?v=DK5OKKbF6GI)
/*
Given a string s, return the longest palindromic substring in s.

Analysis:
- racecar
- xxaabbaayy

_ _ _ _ i _ _ _
0 1 2 3 4 5 6 7
current i = 4

if length = 5 => (s, i, i)
start = 2 = 4 - (5-1)/2
end = 6 = 4 + 5/2

if length = 6 => (s, i, i+1)
start = 2 = 4 - (6-1)/2
end = 7 = 4 + 6/2

Time Complexity: O(n^2)
1. The Outer Loop ($n$): You are correct that the loop for i := range s runs $n$ times.
2. The Inner Work (Not Constant): Inside the loop, you call outwardComparison.
	- In the best case (e.g., string "abcdef"), the expansion stops immediately. This would be $O(1)$ per character, leading to $O(n)$ total.
	- In the worst case (e.g., string "aaaaa..."), for every character i, the outwardComparison function expands all the way to the edges of the string.
	- If i is in the middle of a string of length $n$, the inner loop runs roughly $n/2$ times.
3. Total Complexity: $$n \text{ (outer loop)} \times n \text{ (inner expansion)} = O(n^2)$$
*/
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := range s {
		len1 := outwardComparison(s, i, i)
		len2 := outwardComparison(s, i, i+1)
		maxLen := max(len1, len2)

		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func longestPalindromeExercise(s string) string {
	return ""
}

func outwardComparison(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		// Based on the constraints, s consist of only digits and English letters. So it can just use "s" for iteration directly
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return right - left - 1
}

func outwardComparisonExercise(s string, left, right int) int {
	return 0
}

// LongestPalindromeInefficient preserves the original implementation for reference.
//
// Time Complexity Analysis:
// Although the algorithmic complexity is still O(n^2), the constant factors are much higher here.
// 1. Memory Allocation: Inside the loop, `s[left : right+1]` creates a new string header repeatedly.
// 2. GC Pressure: Creating many temporary string objects triggers the Garbage Collector more often.
// 3. Data Movement: Returning strings involves passing data structures around, whereas passing ints is purely register-based.
func LongestPalindromeInefficient(s string) string {
	if len(s) < 2 {
		return s
	}

	var result string
	for i := range s {
		palindrome1 := outwardComparisonString(s, i, i)
		palindrome2 := outwardComparisonString(s, i, i+1)

		if len(palindrome1) > len(result) {
			result = palindrome1
		}
		if len(palindrome2) > len(result) {
			result = palindrome2
		}
	}

	return result
}

// outwardComparisonString is the helper for the inefficient version.
// It returns the actual substring, causing allocations inside the loop.
func outwardComparisonString(s string, left, right int) string {
	var result string
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		result = s[left : right+1]
		left--
		right++
	}
	return result
}
