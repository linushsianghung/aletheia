package Medium

import (
	"bytes"
	"strconv"
)

// https://leetcode.com/problems/count-and-say/description/
// Ref: A Comprehensive Guide for Golang String Concatenation https://medium.com/@teamcode20233/a-comprehensive-guide-for-golang-string-concatenation-df1a02210492
/*
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.

To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit.
Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.
*/
func countAndSay(n int) string {
	countAndSay := "1"

	for i := 1; i < n; i++ {
		countAndSay = buildCountAndSay(countAndSay)
	}

	return countAndSay
}

func buildCountAndSay(countAndSay string) string {
	var buffer bytes.Buffer
	count, say := 1, countAndSay[0]

	for i := 1; i < len(countAndSay); i++ {
		if countAndSay[i] == say {
			count++
		} else {
			buffer.WriteString(strconv.Itoa(count))
			buffer.WriteString(string(say))
			count = 1
			say = countAndSay[i]
		}
	}
	buffer.WriteString(strconv.Itoa(count))
	buffer.WriteString(string(say))
	return buffer.String()
}
