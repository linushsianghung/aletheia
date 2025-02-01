package BinarySearch

// MySqrt https://leetcode.com/problems/sqrtx/
// Ref:
// - https://www.youtube.com/watch?v=zdMhGxRWutQ
// - https://leetcode.com/problems/sqrtx/solutions/25198/3-java-solutions-with-explanation/
/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

Analysis:
Sample1: 13
left, right := 1, 13: mid = (1 + 13) / 2 = 7 => 7*7 > 13 => right = 7 - 1
left, right := 1, 6: mid = (1 + 6) / 2 = 3.5 => 3*3 < 13 => left = 3 + 1
left, right := 4, 6: mid = (4 + 6) / 2 = 5 => 7*7 > 13 => right = 5 - 1
left, right := 4, 4: mid = (4 + 4) / 2 = 4 => 4*4 > 13 => right = 4 - 1

left = 4, right = 3

Sample2: 17
left, right := 1, 17: mid = (1 + 17) / 2 = 9 => 9*9 > 17 => right = 9 - 1
left, right := 1, 8: mid = (1 + 8) / 2 = 4.5 => 4*4 < 17 => left = 4 + 1
left, right := 5, 8: mid = (5 + 8) / 2 = 6.5 => 6*6 > 17 => right = 6 - 1
left, right := 5, 5: mid = (5 + 5) / 2 = 5 => 5*5 > 17 => right = 5 - 1

left = 5, right = 4
*/
func MySqrt(x int) int {
	if x == 0 {
		return 0
	}

	return mySqrtBS(x)
}

// This is implemented based on the template and more precise
func mySqrtBS(x int) int {
	left, right := 1, x

	for left <= right {
		mid := left + (right-left)/2

		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}

	return right
}

func mySqrtBSExercise(x int) int {
	return 0
}

func mySqrtBSAlt(x int) int {
	left, right := 0, x
	for left < right {
		mid := left + (right-left)/2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func mySqrtNewton(x int) int {
	r := x

	for r*r > x {
		r = (r + x/r) / 2
	}

	return r
}
