package Stack

import "unicode"

// https://leetcode.com/problems/decode-string/description/?envId=leetcode-75
/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.
*/
func decodeString(s string) string {
	result, sRune := "", []rune(s)
	countStack, charStack := make([]int, 0), make([]string, 0)

	for i := 0; i < len(sRune); {
		char := sRune[i]
		switch {
		case unicode.IsLetter(char):
			cRune := make([]rune, 0)
			for unicode.IsLetter(sRune[i]) {
				cRune = append(cRune, sRune[i])
				i++
			}

			charStack = append(charStack, string(sRune[i]))
		case unicode.IsDigit(char):
			var count int
			for unicode.IsDigit(sRune[i]) {
				count = (count * 10) + int(sRune[i]-'0')
				i++
			}

			countStack = append(countStack, count)
		case char == ']':
			count := countStack[len(countStack)-1]
			char := charStack[len(charStack)-1]
			countStack = countStack[:len(countStack)-1]
			charStack = charStack[:len(charStack)-1]

			for i := 1; i < count; i++ {
				result += char
			}
			i++
		}

	}

	return result
}
