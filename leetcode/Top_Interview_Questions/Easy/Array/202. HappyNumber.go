package Array

// https://leetcode.com/problems/happy-number/description/
/*
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:
- Starting with any positive integer, replace the number by the sum of the squares of its digits.
- Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
- Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.
*/
func IsHappy(n int) bool {
	return isHappyMath(n)
	// return isHappyFloydCycle(n)
}

/*
Algorithm:
Using a map to know if the digit has appeared before. If it is true, this will go in to an infinite cycle which means this number would not be happy
*/
func isHappyMath(n int) bool {
	pairs := make(map[int]bool)

	for {
		sum := compute(n)

		if sum == 1 {
			return true
		}
		if ok := pairs[sum]; ok {
			return false
		}

		pairs[sum] = true
		n = sum
	}
}

/*
Analysis:
It was clearly mentioned in the problem statement that if a number ain't happy, then it will lead to a cycle.
Whenever we hear the word "cycle", the first thing we should think of is "Floyd's cycle-finding algorithm" also known as "Tortoise and the Hare algorithm."

Ref: https://www.youtube.com/watch?v=SJRaMCSgNWQ
*/
// <editor-fold desc="Floyd's cycle-finding algorithm">
func isHappyFloydCycle(n int) bool {
	slow, fast := n, n
	for fast != 1 {
		slow = compute(slow)
		fast = compute(compute(fast))

		if slow == fast && fast != 1 {
			return false
		}
	}

	return true
}

// </editor-fold>

func compute(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
