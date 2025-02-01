package Stack

// https://leetcode.com/problems/backspace-string-compare
// Reference: https://leetcode.com/problems/backspace-string-compare/solutions/570511/c-simple-and-easy-explanation-100-memory-and-100-speed-0ms-o-1-space-o-n-time/
/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
*/
func backspaceCompare(s string, t string) bool {
	return backspaceCompare2Stack(s, t)
	//return backspaceCompare2Pointers(s, t)
}

func backspaceCompare2Stack(s string, t string) bool {
	stackS, stackT := make([]rune, 0), make([]rune, 0)

	for _, c := range s {
		if len(stackS) > 0 && c == '#' {
			stackS = stackS[:len(stackS)-1]
			continue
		} else if c != '#' {
			stackS = append(stackS, c)
		}
	}

	for _, c := range t {
		if len(stackT) > 0 && c == '#' {
			stackT = stackT[:len(stackT)-1]
			continue
		} else if c != '#' {
			stackT = append(stackT, c)
		}
	}

	if len(stackS) != len(stackT) {
		return false
	}

	for i, s := range stackS {
		if stackT[i] != s {
			return false
		}
	}

	return true
}

func backspaceCompare2StackExercise(s string, t string) bool {

	return false
}

func backspaceCompare2Pointers(s string, t string) bool {
	sharpNum := 0
	tailS, tailT := len(s)-1, len(t)-1
	for {
		for ; tailS >= 0; tailS-- {
			if s[tailS] == '#' {
				sharpNum++
			} else {
				if sharpNum > 0 {
					sharpNum--
				} else {
					break
				}
			}
		}

		sharpNum = 0
		for ; tailT >= 0; tailT-- {
			if t[tailT] == '#' {
				sharpNum++
			} else {
				if sharpNum > 0 {
					sharpNum--
				} else {
					break
				}
			}
		}

		if tailS >= 0 && tailT >= 0 && s[tailS] == t[tailT] {
			tailS--
			tailT--
			continue
		}
		break
	}
	return tailS == -1 && tailT == -1
}

func backspaceCompare2PointersExercise(s string, t string) bool {

	return false
}
