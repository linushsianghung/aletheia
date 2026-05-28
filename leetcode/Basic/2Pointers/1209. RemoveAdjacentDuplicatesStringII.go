package TwoPointers

import "github.com/linushung/aletheia/leetcode/Basic/Stack"

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
/*
You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them,
causing the left and the right side of the deleted substring to concatenate together.

We repeatedly make k duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.
*/
func removeDuplicatesII(s string, k int) string {
	return removeDuplicatesII2Pointers(s, k)
}

func removeDuplicatesII2Pointers(s string, k int) string {
	// 'anchor' acts as the stack pointer. It points to the index where the next character will be written.
	anchor := 0
	// 'counts' stores the consecutive frequency of the character at the corresponding index in 'stack'
	counts := make([]int, len(s))
	// Convert to rune slice to handle potential multi-byte characters and allow in-place modification
	stack := []rune(s)

	for _, r := range stack {
		// Push character onto the "stack"
		stack[anchor] = r

		// Calculate count: If it matches the previous char in our stack, increment count
		if anchor > 0 && stack[anchor-1] == r {
			counts[anchor] = counts[anchor-1] + 1
		} else {
			counts[anchor] = 1
		}

		// Check if we reached k duplicates. If so, "pop" them by moving the anchor pointer back.
		if counts[anchor] == k {
			anchor -= k
		}
		anchor++
	}

	return string(stack[:anchor])
}

func removeDuplicatesIIExercise(s string, k int) string {
	return ""
}

// Related Problem: 1047. Remove All Adjacent Duplicates In String: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
func removeDuplicates1047(s string) {
	Stack.RemoveDuplicates(s)
}

/*
# Analysis:
1. Concept Introduction:
   Imagine you are playing a game like Candy Crush or Zuma. When a certain number of identical items (k) align,
   they vanish, and the items that were on either side "collapse" together. This collapse can trigger a chain
   reaction if the newly touching items also form a group of k.

   To solve this efficiently, we use a "Stack" logic. Instead of actually deleting parts of a string (which is
   expensive O(n) per deletion), we use a pointer (`anchor`) to keep track of our "current" end of the string.
   We also maintain a `counts` array. This array acts as a history: it tells us, "at this specific position in our
   newly built string, how many identical characters have we seen consecutively?"

2. Complexity Analysis:
   - Time Complexity: O(n), where n is the length of the string s. We iterate through the string exactly once.
     Each character is added to the "stack" once and potentially removed once.
   - Space Complexity: O(n). We use an integer array `counts` of size n and a character array `stack` of size n.
     In Java, `s.toCharArray()` creates a new array, and the final `new String(...)` creates another, but the
     peak auxiliary space is proportional to n.

3. How to optimize step-by-step during an Interview:
   - Step 1 (Brute Force): Suggest finding k-duplicates, deleting them, and restarting. Explain why this is
     slow (O(n^2 / k)) due to repeated string scanning and slicing.
   - Step 2 (Stack of Pairs): Suggest using a `Stack<Pair<Character, Integer>>`. This handles the "collapse"
     logic cleanly. If the top of the stack matches the current char, increment the count; if the count reaches k,
     pop.
   - Step 3 (Two Pointers / Array as Stack): Realize that the stack of objects has overhead. Optimize by using
     two arrays (one for chars, one for counts) and a pointer to simulate the stack. This is the most performant
     approach in both Go and Java.

Implementation Method:
public String removeDuplicates(String s, int k) {
	int anchor = 0;
	int[] counts = new int[s.length()];
	char[] stack = s.toCharArray(); // Convert to char array for "in-place" simulation

	for (var c : s.toCharArray()) {
		stack[anchor] = c;

		if (anchor > 0 && stack[anchor-1] == c) {
			count[anchor] = count[anchor-1] + 1;
		} else {
			count[anchor] = 1;
		}

		if (count[anchor] == k) {
			anchor -= k;
		}
		anchor++;
	}

	return new String(stack, 0, anchor);
}
*/
