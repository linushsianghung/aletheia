package Basic

import "unicode"

// https://leetcode.com/problems/to-lower-lcase/
/* Given a string s, return the string after replacing every uppercase letter with the same lowercase letter. */
func toLowerCase(str string) string {
	if str == "" {
		return ""
	}

	lower := ""
	for _, s := range str {
		if unicode.IsUpper(s) {
			lower += string(s + 32)
		} else {
			lower += string(s)
		}

	}
	return lower
}
