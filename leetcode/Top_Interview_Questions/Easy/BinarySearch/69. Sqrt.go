package BinarySearch

// https://leetcode.com/problems/sqrtx/
// Ref:
// - https://www.youtube.com/watch?v=zdMhGxRWutQ
// - https://leetcode.com/problems/sqrtx/solutions/25198/3-java-solutions-with-explanation/
/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.
*/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	return mySqrtBS(x)
}

func mySqrtBS(x int) int {
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
