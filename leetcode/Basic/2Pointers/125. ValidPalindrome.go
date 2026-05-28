package TwoPointers

import (
	"strings"
	"unicode"
)

// IsPalindrome https://leetcode.com/problems/valid-palindrome/
/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)
	sRune := []rune(s)

	left, right := 0, len(s)-1
	for left < right {
		if !unicode.IsLetter(sRune[left]) && !unicode.IsDigit(sRune[left]) {
			left++
			continue
		}
		if !unicode.IsLetter(sRune[right]) && !unicode.IsDigit(sRune[right]) {
			right--
			continue
		}

		if sRune[left] != sRune[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isPalindromeExercise(s string) bool {
	return false
}

/*
# Analysis
1. Concept Introduction:
The Two-Pointer technique is the bread and butter of string/array manipulation. Imagine two people standing at opposite ends of a hallway walking toward each other. They only stop to look at "important" items (alphanumeric chars). If they ever see different items at the same time, the "symmetry" (palindrome property) is broken.

2. Complexity Analysis:
- Time: O(N) -> We visit each character at most once.
- Space: O(1) -> We don't allocate any extra data structures that grow with input size. We use `s.charAt(i)` to avoid creating a new char array.

3. Interview Suggestion:
- Clarify Requirements: Ask if "alphanumeric" includes just `a-z`, `A-Z`, and `0-9`. (It usually does).
- Standard Library usage: In Java, mention `Character.isLetterOrDigit()` and `Character.toLowerCase()`. It shows you know the language's standard API.
- Edge Cases: Always mention empty strings or strings with only special characters (e.g., ".,"). In these cases, the pointers will cross, and the function should return `true`.
- Optimization Talk: Mention that while `s.toLowerCase()` is easier to write, doing it character-by-character inside the loop is more memory-efficient as it avoids creating a whole new string object in the heap.

Implementation Method:
public boolean isPalindrome(String s) {
	if (s == null || s.isEmpty()) {
		return true;
	}

	int left = 0, right = s.length() - 1;

	while (left < right) {
		if (!Character.isLetterOrDigit(s.charAt(left))) {
			left++;
			continue;
		}
		if (!Character.isLetterOrDigit(s.charAt(right))) {
			right--;
			continue;
		}
		// Both are alphanumeric, compare them
		if (Character.toLowerCase(s.charAt(left)) != Character.toLowerCase(s.charAt(right))) {
			return false;
		}
		left++;
		right--;
	}

	return true;
}
*/
