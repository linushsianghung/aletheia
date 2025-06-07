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

func removeDuplicates(s string, k int) string {
	count, stack := 0, make([]rune, 0)

	for _, r := range s {
		if len(stack) > 0 && stack[len(stack)-1] == r {
			count++
			if count == k {
				stack = stack[:len(stack)-k+1]
				count++
				continue
			}
		}

		stack = append(stack, r)
	}
	return string(stack)
}

func removeDuplicatesII2Pointers(s string, k int) string {
	anchor := 0
	// Using a count slice to check if each character is k adjacent letter
	sRune, count := []rune(s), make([]int, len(s))

	for _, r := range sRune {
		sRune[anchor] = r
		if anchor > 0 && sRune[anchor-1] == r {
			count[anchor] = count[anchor-1] + 1
		} else {
			count[anchor] = 1
		}
		if count[anchor] == k {
			anchor -= k
		}
		anchor++
	}

	return string(sRune[:anchor])
}

func removeDuplicatesII2PointersExercise(s string, k int) string {
	return ""
}

// Related Problem: 1047. Remove All Adjacent Duplicates In String: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
func removeDuplicates1047(s string) {
	Stack.RemoveDuplicates(s)
}
