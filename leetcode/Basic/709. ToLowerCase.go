package Basic

import "unicode"

// https://leetcode.com/problems/to-lower-case/
/* Given a string s, return the string after replacing every uppercase letter with the same lowercase letter. */
func toLowerCase(s string) string {
	if s == "" {
		return ""
	}

	//lower := ""
	//for _, s := range s {
	//	if unicode.IsUpper(s) {
	//		lower += string(s + 32)
	//	} else {
	//		lower += string(s)
	//	}
	//
	//}
	//return lower

	result := make([]rune, 0)

	for _, r := range s {
		if unicode.IsUpper(r) {
			r = r + 32
		}
		result = append(result, r)
	}
	return string(result)
}

func toLowerCaseExercise(str string) string {
	return ""
}
