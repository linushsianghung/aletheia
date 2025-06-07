package Array_String

// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envId=leetcode-75
// Reference: Euclidean Algorithm: https://jason-chen-1992.weebly.com/home/-euclidean-algorithm#google_vignette
/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Reference: https://leetcode.com/problems/greatest-common-divisor-of-strings/solutions/3124940/c-one-line-beats-100-runtime-explanation
Analysis:
- str1+str2 == str2+str1 if and only if str1 and str2 have a gcd. E.g. str1 = abcabc, str2 = abc, then str1+str2 = abcabcabc = str2+str1
- This(str1+str2 == str2+str1) is a requirement for the strings to have a gcd. If one of them is NOT a common part then gcd is "". It means we will return empty string
- Proof: str1 = mGCD, str2 = nGCD, str1 + str2 = (m + n)GCD = str2 + str1 (Both strings are made of same substring added multiple times.)
- Since they are multiples, next step is simply to find the gcd of the lengths of 2 strings e.g. gcd(6,3) = 3. We can use gcd function to find that and get the substring of length 3 from either str1 or str2.
*/
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	var gcd func(p, q int) int
	gcd = func(p, q int) int {
		if q == 0 {
			return p
		}
		return gcd(q, p%q)
	}

	length := gcd(len(str1), len(str2))
	return str1[:length]
}
