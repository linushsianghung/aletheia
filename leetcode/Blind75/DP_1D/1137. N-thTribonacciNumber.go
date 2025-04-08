package DP_1D

// https://leetcode.com/problems/n-th-tribonacci-number/description/?envId=leetcode-75
/*
The Tribonacci sequence Tn is defined as follows:
T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.
*/
func tribonacci(n int) int {
	return tribonacciHelper(n, make(map[int]int))
}

func tribonacciHelper(n int, note map[int]int) int {
	if value, ok := note[n]; ok {
		return value
	}

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	note[n] = tribonacciHelper(n-1, note) + tribonacciHelper(n-2, note) + tribonacciHelper(n-3, note)
	return note[n]
}
