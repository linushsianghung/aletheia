package Basic

// IsValid https://leetcode.com/problems/valid-parentheses/
/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
func IsValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		switch c {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		default:
			if len(stack) == 0 || c != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}

func isValidExercise(s string) bool {

	return false
}
