package testdome

import (
	"fmt"
	"math"
)

/**
Implement the function findRoots to find the roots of the quadratic equation: ax2 + bx + c = 0. If the equation has only
one solution, the function should return that solution as both results. The equation will always have at least one solution.

The roots of the quadratic equation can be found with the following formula: A quadratic equation.

For example, the roots of the equation 2x2 + 10x + 8 = 0 are -1 and -4.
*/
func findRoots(a, b, c float64) (float64, float64) {

	deter := math.Sqrt(b*b - 4*a*c)

	if deter > 0 {
		return (-b + deter) / (2 * a), (-b - deter) / (2 * a)
	} else if deter == 0 {
		root := (-b) / (2 * a)
		return root, root
	}

	return 0, 0
}

func QuadraticEquation() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f\n", x1, x2)
}
