package Stack

// https://leetcode.com/problems/make-the-string-great
/*
Given a string s of lower and upper case English letters.

A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:

0 <= i <= s.length - 2
s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice versa.
To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.

Return the string after making it good. The answer is guaranteed to be unique under the given constraints.

Notice that an empty string is also good.
*/
func makeGood(s string) string {
	stack := make([]rune, 0)

	for _, r := range s {
		if len(stack) > 0 && (stack[len(stack)-1]-r == 32 || r-stack[len(stack)-1] == 32) {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, r)
	}

	return string(stack)
}

func makeGoodExercise(s string) string {
	return ""
}

// Related Problem: 1047. Remove All Adjacent Duplicates In String: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
func removeDuplicatesFrom1047(s string) {
	RemoveDuplicates(s)
}

/*
Analysis:
1. Concept:
   The problem requires us to remove adjacent "bad" pairs. A pair is bad if they are the same letter but different cases.
   In ASCII/Unicode, the difference between 'a' (97) and 'A' (65) is exactly 32.
   A Stack is the ideal data structure here because removing a pair might make the characters that were previously
   separated become adjacent, potentially forming a new bad pair (e.g., "abBA" -> remove "bB" -> "aA" -> remove "aA" -> "").

2. Complexity Analysis:
   - Time Complexity: O(n), where n is the length of the string. We iterate through the string once, and each
     character is pushed to or popped from the stack at most once.
   - Space Complexity: O(n) in the worst case (e.g., "abcdef") where no characters are removed and we store
     the entire string in the stack/StringBuilder.

3. Interview Suggestion:
   - Mention the ASCII property: explain that `Math.abs(charA - charB) == 32` is a quick way to check for same-letter-different-case.
   - String Mutability: In Java interviews, always highlight why you chose `StringBuilder` over `String` concatenation to avoid O(n^2) time.
   - Edge Cases: Mention handling of empty strings or strings with only one character.

Implementation Method:
public String makeGood(String s) {
    StringBuilder sb = new StringBuilder();

    for (char c : s.toCharArray()) {
        int n = sb.length();
        if (n > 0) {
            char lastChar = sb.charAt(n - 1);
            // Check if characters are the same letter with different cases
            if (Math.abs(lastChar - c) == 32) {
                sb.deleteCharAt(n - 1);
                continue;
            }
        }
        sb.append(c);
    }

    return sb.toString();
}
*/
