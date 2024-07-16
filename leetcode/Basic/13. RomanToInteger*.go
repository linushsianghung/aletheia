package Basic

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := symbols[s[len(s)-1]]
	for i := len(s) - 1; i > 0; i-- {
		if symbols[s[i-1]] >= symbols[s[i]] {
			total += symbols[s[i-1]]
		} else {
			total -= symbols[s[i-1]]
		}
	}

	return total
}
