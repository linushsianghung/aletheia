package Backtracking

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/?envId=leetcode-75
// Reference: https://www.youtube.com/watch?v=a-sMgZ7HGW0
/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	// It's too difficult to convert unicode (derived from the element of string slice) to integer
	// digitMapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	digitMapping := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	return backtrackLetterCombination(digits, digitMapping)
}

func backtrackLetterCombination(digits string, mapping map[rune]string) []string {
	result := make([]string, 0)

	var nestedFunc func(processor []rune, index int)
	nestedFunc = func(processor []rune, index int) {
		if len(processor) == len(digits) {
			result = append(result, string(processor))
			return
		}

		for _, c := range mapping[rune(digits[index])] {
			processor = append(processor, c)
			nestedFunc(processor, index+1)
			processor = processor[:len(processor)-1]
		}
	}

	nestedFunc(make([]rune, 0), 0)
	return result
}

func backtrackLetterCombinationAlt(result *[]string, digits, curStr string, index int, mapping map[rune]string) {
	if len(curStr) == len(digits) {
		*result = append(*result, curStr)
		return
	}

	for _, c := range mapping[rune(digits[index])] {
		backtrackLetterCombinationAlt(result, digits, curStr+string(c), index+1, mapping)
	}
}
