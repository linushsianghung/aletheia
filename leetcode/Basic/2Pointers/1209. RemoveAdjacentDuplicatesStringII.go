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
	// Convert to rune slice to handle potential multi-byte characters and allow in-place modification
	stack := []rune(s)
	// 'counts' stores the consecutive frequency of the character at the corresponding index in 'stack'
	counts := make([]int, len(s))

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
