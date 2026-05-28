package Basic

import "math"

// Reverse https://leetcode.com/problems/reverse-integer/
/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
func Reverse(x int) int {
	rev := 0
	// It has to be non-equal because of negative cases, like -123 to -321
	for x != 0 {
		rev = rev*10 + x%10
		if rev > math.MaxInt32 || rev < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return rev
}

func reverseExercise(x int) int {
	return 0
}

/*
Analysis:
1. Concept:
   The core idea is to "pop" the last digit of the integer using the modulo operator (`x % 10`)
   and "push" it into a new variable (`rev`) by multiplying the existing result by 10.
   Analogy: Imagine a stack of digits. We take from the top (the units place) and build a
   new stack from the bottom up.

2. Complexity Analysis:
   - Time Complexity: O(log10(x)). We iterate through the digits of x. The number of digits
     in x is roughly log10(x).
   - Space Complexity: O(1). We only use a few variables regardless of the input size.

3. Optimized Solution Approach:
   To solve this during an interview without using 64-bit integers:
   - Step 1: Initialize `rev = 0`.
   - Step 2: Loop while `x != 0`.
   - Step 3: Identify the "pop" digit: `pop = x % 10`.
   - Step 4: **Crucial Step (Overflow Check)**: Before updating `rev`, check if `rev * 10 + pop`
     will exceed 32-bit limits.
     - For positive: If `rev > Max/10`, or `rev == Max/10` and `pop > 7` (since 2,147,483,647 ends in 7).
     - For negative: If `rev < Min/10`, or `rev == Min/10` and `pop < -8` (since -2,147,483,648 ends in 8).
   - Step 5: Update `rev = rev * 10 + pop` and `x /= 10`.

Implementation Method (Java):
public int reverse(int x) {
    int rev = 0;
    while (x != 0) {
        int pop = x % 10;
        x /= 10;

        // Check for overflow before it happens
        if (rev > Integer.MAX_VALUE/10 || (rev == Integer.MAX_VALUE / 10 && pop > 7)) return 0;
        if (rev < Integer.MIN_VALUE/10 || (rev == Integer.MIN_VALUE / 10 && pop < -8)) return 0;

        rev = rev * 10 + pop;
    }
    return rev;
}
*/
