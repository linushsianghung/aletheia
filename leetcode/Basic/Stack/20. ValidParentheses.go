package Stack

// IsValid https://leetcode.com/problems/valid-parentheses/
/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
func IsValid(s string) bool {
	stack := make([]rune, 0)

	for _, r := range s {
		switch r {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != r {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValidExercise(s string) bool {

	return false
}
