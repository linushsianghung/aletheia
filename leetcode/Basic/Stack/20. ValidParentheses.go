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

/*
Analysis:
1. Concept:
   The problem follows a "Last-In-First-Out" (LIFO) pattern. When we encounter an opening bracket,
   it must be closed by the most recent corresponding closing bracket. This is a classic
   application for a Stack. Think of it like a stack of cafeteria trays; you can only
   meaningfully interact with the one on top.

2. How to come up with the optimized solution:
   - Step 1: Realize that nested structures require memory of "what is currently open."
   - Step 2: A simple counter won't work because there are three different types of brackets
     that must close in a specific order (e.g., "([)]" is invalid even if counts match).
   - Step 3: Use a Stack. To make the code cleaner, when you see an opening bracket, push
     its *required* closing partner onto the stack. This way, when you encounter a closing
     bracket in the input, you just need to check if it matches the `pop()` result.

3. Complexity Analysis:
   - Time Complexity: O(n), where n is the length of the string. We traverse the string exactly once.
   - Space Complexity: O(n). In the worst-case scenario (e.g., "((((("), we push all characters
     onto the stack.

Implementation Method:
public boolean isValid(String s) {
	// Using the legacy Stack class for simplicity, although Deque<Character> stack = new ArrayDeque<>() is often preferred in production.
	Stack<Character> stack = new Stack<>();

	for (char c : s.toCharArray()) {
		if (c == '(') {
			stack.push(')');
		} else if (c == '{') {
			stack.push('}');
		} else if (c == '[') {
			stack.push(']');
		} else {
			// If stack is empty, it means we have a closing bracket without an opening one.
			// If stack.pop() != c, it means the brackets are closing in the wrong order.
			if (stack.isEmpty() || stack.pop() != c) {
				return false;
			}
		}
	}

	// If the stack is empty, all brackets were matched correctly.
	return stack.isEmpty();
}
*/
