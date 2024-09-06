package Stack

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
/*
You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.

We repeatedly make duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.
*/
func removeDuplicates(s string) string {
	//return removeDuplicatesStack(s)
	return removeDuplicates2Pointers(s)
}

func removeDuplicates2Pointers(s string) string {
	anchor := 0
	sRune := []rune(s)

	for _, r := range sRune {
		sRune[anchor] = r
		if anchor > 0 && sRune[anchor-1] == r {
			anchor -= 2
		}
		anchor++
	}

	return string(sRune[:anchor])
}

func removeDuplicatesStack(s string) string {
	stack := make([]rune, 0)

	for _, r := range s {
		if len(stack) > 0 && stack[len(stack)-1] == r {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, r)
	}

	return string(stack)
}
