package Stack

// https://leetcode.com/problems/removing-stars-from-a-string/description/?envId=leetcode-75
/*
You are given a string s, which contains stars *.

In one operation, you can:
- Choose a star in s.
- Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:
The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.
*/
func removeStars(s string) string {
	//return removeStars2Pointers(s)
	return removeStarsStack(s)
}

func removeStarsStack(s string) string {
	strRune, stack := []rune(s), make([]rune, 0)

	for _, r := range strRune {
		if r == '*' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, r)
	}

	result := string(stack)
	return result
}

func removeStars2Pointers(s string) string {
	slow, strRune := 0, []rune(s)

	for fast := 0; fast < len(s); fast++ {
		if strRune[fast] == '*' {
			slow--
			continue
		}

		strRune[slow] = strRune[fast]
		slow++
	}

	return string(strRune[:slow])
}
