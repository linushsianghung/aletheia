package Array

import "strconv"

// https://leetcode.com/problems/fizz-buzz/description/
/*
Given an integer n, return a string array answer (1-indexed) where:
- answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
- answer[i] == "Fizz" if i is divisible by 3.
- answer[i] == "Buzz" if i is divisible by 5.
- answer[i] == i (as a string) if none of the above conditions are true.
*/

func fizzBuzz(n int) []string {
	answer := make([]string, 0)

	// for num := range n: Cannot use range n because Go will iterate from 0 to one less than n
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			answer = append(answer, "FizzBuzz")
		case i%5 == 0:
			answer = append(answer, "Buzz")
		case i%3 == 0:
			answer = append(answer, "Fizz")
		default:
			answer = append(answer, strconv.Itoa(i))
		}
	}

	return answer
}

func fizzBuzzExercise(n int) []string {
	return nil
}
