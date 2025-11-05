package Stack

// https://leetcode.com/problems/backspace-string-compare
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

	for _, r := range s {
		if r == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
			continue
		}

		stackS = append(stackS, r)
	}

	for _, r := range t {
		if r == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
			continue
		}

		stackT = append(stackT, r)
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
