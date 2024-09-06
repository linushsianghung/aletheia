package Stack

import "math"

// https://leetcode.com/problems/asteroid-collision/description/?envId=leetcode-75
/*
We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left).
Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.
*/
func asteroidCollision(asteroids []int) []int {
	stack := []int{asteroids[0]}

	for i := 1; i < len(asteroids); {
		if len(stack) == 0 {
			stack = append(stack, asteroids[i])
			i++
			continue
		}
		// There are totally 4 scenarios will happen and only 3 needs to be handle, for other scenarios just pushes asteroids (back) to the stack
		// 1. + +
		// 2. - -
		// 3. + -
		// 4. - +
		current := stack[len(stack)-1]
		if current < 0 || (current > 0 && asteroids[i] > 0) {
			stack = append(stack, asteroids[i])
			i++
			continue
		}

		// Handling scenario 3
		if math.Abs(float64(current)) > math.Abs(float64(asteroids[i])) {
			i++
		} else if math.Abs(float64(current)) < math.Abs(float64(asteroids[i])) {
			stack = stack[:len(stack)-1]
		} else {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return stack
}
