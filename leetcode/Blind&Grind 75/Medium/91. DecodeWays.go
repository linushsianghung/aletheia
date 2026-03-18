package Medium

import (
	"github.com/linushung/aletheia/leetcode/Top_Interview_Questions/Easy/DynamicProgramming"
)

// NumDecodings https://leetcode.com/problems/decode-ways/
/*
You have intercepted a secret message encoded as a string of numbers. The message is decoded via the following mapping:

"1" -> 'A'

"2" -> 'B'

...

"25" -> 'Y'

"26" -> 'Z'

However, while decoding the message, you realize that there are many different ways you can decode the message because some codes are contained in other codes ("2" and "5" vs "25").

For example, "11106" can be decoded into:

"AAJF" with the grouping (1, 1, 10, 6)
"KJF" with the grouping (11, 10, 6)
The grouping (1, 11, 06) is invalid because "06" is not a valid code (only "6" is valid).
Note: there may be strings that are impossible to decode.

Given a string s containing only digits, return the number of ways to decode it. If the entire string cannot be decoded in any valid way, return 0.

The test cases are generated so that the answer fits in a 32-bit integer.

Analysis:
Think of this problem like a decision tree or "Climbing Stairs": At any given digit in the string, you have at most two choices:
1.Take 1 Digit:
	- You can decode the single digit s[i] if it is between '1' and '9'.
	- (You cannot decode '0' by itself).
	- Result: Move to index i+1.
2. Take 2 Digits:
	- You can decode the pair s[i]s[i+1] if it forms a number between 10 and 26.
	- Result: Move to index i+2.

The total number of ways to decode from index i is simply: (text{Ways from } i+1) + (text{Ways from } i+2)


      dfs(0) for "1111"
     /                \
  dfs(1) for "111"    dfs(2) for "11"  <-- We need this result
 /        \
dfs(2) for "11"  dfs(3) for "1"
  ^
  |-- Whoops! We are calculating this AGAIN!
*/
func NumDecodings(s string) int {
	return numDecodingsMemo(s)
}

// Improvements of numDecodingsAlternative:
// 1. Uses 'index' (int) as the map key instead of substring to save memory and time.
// 2. Uses character literals ('0', '1') instead of ASCII codes (48, 49) for readability.
func numDecodingsMemo(s string) int {
	// memo := make(map[int]int)

	// Optimization: Use a slice initialized to -1. 0 is a valid result (cannot decode), so we can't use 0 as "empty".
	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}

	var decodeFunc func(index int) int
	decodeFunc = func(index int) int {
		// Base Case: Reached end of string successfully
		if index == len(s) {
			return 1
		}
		// Base Case: Leading zero is invalid
		if s[index] == '0' {
			return 0
		}

		// Memoization Check: If we have already computed the result for this index, return the cached value immediately to avoid re-computing. This is the core of the dynamic programming optimization.
		if memo[index] != -1 {
			return memo[index]
		}

		// 1. Decode single digit
		count := decodeFunc(index + 1)
		// 2. Decode two digits (must be between "10" and "26")
		if index < len(s)-1 {
			if s[index] == '1' || (s[index] == '2' && s[index+1] <= '6') {
				count += decodeFunc(index + 2)
			}
		}

		memo[index] = count
		return count
	}

	return decodeFunc(0)
}

// numDecodingsTable is the Iterative (Bottom-Up) approach.
// It avoids recursion depth issues and is often preferred for production code.
func numDecodingsTable(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	// dp[i] represents the number of ways to decode the string of length i
	dp := make([]int, len(s)+1)
	dp[0] = 1 // Base case for length 0
	dp[1] = 1 // Base case for length 1 (we already checked s[0] != '0')

	for i := 2; i <= len(s); i++ {
		// Check if single digit (s[i-1]) is valid (1-9)
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// Check if two digits (s[i-2:i]) are valid (10-26)
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func numDecodingsExercise(s string) int {
	return 0
}

/*
This implementation is hard to read because:
1. Magic Numbers: Using 48, 49, 50 (ASCII codes) forces the reader to mentally translate them to '0', '1', '2'.
2. Inefficient Map Key: It uses s[index:] (the substring) as the map key. This creates a new string in memory for every recursive call. The index (integer) is sufficient to identify the state.
*/
func numDecodingsAlternative(s string) int {
	var decodingFunc func(s string, index int, note map[string]int) int
	decodingFunc = func(s string, index int, note map[string]int) int {
		if index > len(s) {
			return 0
		}
		if index == len(s) {
			return 1
		}
		if s[index] == 48 {
			return 0
		}

		if val, ok := note[s[index:]]; ok {
			return val
		}

		count := decodingFunc(s, index+1, note)
		if index < len(s)-1 && (s[index] == 49 || (s[index] == 50 && s[index+1] <= 54)) {
			count += decodingFunc(s, index+2, note)
		}
		note[s[index:]] = count

		return note[s[index:]]
	}

	return decodingFunc(s, 0, make(map[string]int))
}

// Related Problem: 70. Climbing Stairs: https://leetcode.com/problems/climbing-stairs/description/
func climbStairs(n int) int {
	return DynamicProgramming.ClimbStairs(n)
}
