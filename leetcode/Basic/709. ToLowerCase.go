package Basic

import "unicode"

// https://leetcode.com/problems/to-lower-case/
/* Given a string s, return the string after replacing every uppercase letter with the same lowercase letter. */
func toLowerCase(s string) string {
	if s == "" {
		return ""
	}

	result := make([]rune, 0)

	for _, r := range s {
		if unicode.IsUpper(r) {
			r += 32
		}
		result = append(result, r)
	}
	return string(result)
}

func toLowerCaseExercise(str string) string {
	return ""
}
