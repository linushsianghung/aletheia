package Medium

import (
	"slices"
	"strconv"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.
*/
func evalRPN(tokens []string) int {
	stack, operators := make([]int, 0), []string{"+", "-", "*", "/"}

	for _, token := range tokens {
		if slices.Contains(operators, token) {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := calculator(left, right, token)
			stack = append(stack, result)
			continue
		}

		cInt, _ := strconv.Atoi(token)
		stack = append(stack, cInt)
	}

	return stack[0]
}

func calculator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return 0
}
