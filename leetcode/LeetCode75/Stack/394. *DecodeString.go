package Stack

import (
	"strconv"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/decode-string/description/?envId=leetcode-75
/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.
*/
func decodeString(s string) string {
	stack := make([]string, 0)
	num, str := 0, ""

	for _, c := range s {
		switch {
		case c == '[':
			stack = append(stack, str)
			stack = append(stack, strconv.Itoa(num))
			str = ""
			num = 0
		case c == ']':
			preNum, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			preStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			str = preStr + strings.Repeat(str, preNum)
		case unicode.IsDigit(rune(c)):
			num = num*10 + int(rune(c)-'0')
		default:
			str += string(c)
		}
	}

	return str
}
