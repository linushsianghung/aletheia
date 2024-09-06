package Stack

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
	anchor := 0
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
