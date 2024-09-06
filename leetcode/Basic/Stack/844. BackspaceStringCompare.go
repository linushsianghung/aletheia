package Stack

// https://leetcode.com/problems/backspace-string-compare
// Reference: https://leetcode.com/problems/backspace-string-compare/solutions/570511/c-simple-and-easy-explanation-100-memory-and-100-speed-0ms-o-1-space-o-n-time/
/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
*/
func backspaceCompare(s string, t string) bool {
	//return backspaceCompare2Stack(s, t)
	return backspaceCompare2Pointers(s, t)
}

func backspaceCompare2Pointers(s string, t string) bool {
	sharpNum := 0
	tails, tailt := len(s)-1, len(t)-1
	for {
		for ; tails >= 0; tails-- {
			if s[tails] == '#' {
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
		for ; tailt >= 0; tailt-- {
			if t[tailt] == '#' {
				sharpNum++
			} else {
				if sharpNum > 0 {
					sharpNum--
				} else {
					break
				}
			}
		}

		if tails >= 0 && tailt >= 0 && s[tails] == t[tailt] {
			tails--
			tailt--
			continue
		}
		break
	}
	return tails == -1 && tailt == -1
}

func backspaceCompare2Stack(s string, t string) bool {
	stackS, stackT := make([]rune, 0), make([]rune, 0)

	for _, r := range s {
		if r == '#' && len(stackS) > 0 {
			stackS = stackS[:len(stackS)-1]
			continue
		} else if r != '#' {
			stackS = append(stackS, r)
		}
	}

	for _, r := range t {
		if r == '#' && len(stackT) > 0 {
			stackT = stackT[:len(stackT)-1]
			continue
		} else if r != '#' {
			stackT = append(stackT, r)
		}
	}

	if len(stackS) != len(stackT) {
		return false
	}

	for len(stackS) != 0 {
		if stackS[len(stackS)-1] != stackT[len(stackT)-1] {

			return false
		}
		stackS = stackS[:len(stackS)-1]
		stackT = stackT[:len(stackT)-1]
	}

	return true
}
