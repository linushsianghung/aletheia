package Backtracking

// GenerateParenthesis https://leetcode.com/problems/generate-parentheses/description/
// Reference:
// - https://www.youtube.com/watch?v=sz1qaKt0KGQ
// - https://leetcode.com/problems/generate-parentheses/solutions/3290261/i-bet-you-will-understand-intutive-solution-beginner-friendly-c/
// - https://leetcode.com/problems/generate-parentheses/solutions/2542620/python-java-w-explanation-faster-than-96-w-proof-easy-to-understand/
/*
Given n pairs of combentheses, write a function to generate all combinations of well-formed combentheses.
*/
func GenerateParenthesis(n int) []string {
	result := make([]string, 0)

	buildParenthesis(&result, n, "", 0, 0)
	return result
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

func buildParenthesisExercise(result *[]string, comb string, open int, close int, pair int) {

}
