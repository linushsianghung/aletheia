package Backtracking

// GenerateParenthesis https://leetcode.com/problems/generate-parentheses/description/
// Reference:
// - https://www.youtube.com/watch?v=sz1qaKt0KGQ
// - https://leetcode.com/problems/generate-parentheses/solutions/2542620/python-java-w-explanation-faster-than-96-w-proof-easy-to-understand/
/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/
func GenerateParenthesis(n int) []string {
	return buildParenthesisRecursively(n)
}

func buildParenthesisRecursively(pair int) []string {
	result := make([]string, 0)

	var nestedFunc func(parenthesis string, open, close int)
	nestedFunc = func(parenthesis string, open, close int) {
		if len(parenthesis) == pair*2 {
			result = append(result, parenthesis)
			return
		}

		if open < pair {
			nestedFunc(parenthesis+"(", open+1, close)
		}
		if close < open {
			nestedFunc(parenthesis+")", open, close+1)
		}
	}

	nestedFunc("", 0, 0)
	return result
}

func buildParenthesisExercise(pair int) []string {
	return nil
}

func buildParenthesis(result *[]string, pair int, comb string, open int, close int) {
	if len(comb) == pair*2 {
		*result = append(*result, comb)
		return
	}

	if open < pair {
		buildParenthesis(result, pair, comb+"(", open+1, close)
	}
	if close < open {
		buildParenthesis(result, pair, comb+")", open, close+1)
	}
}
