package DynamicProgramming

// ClimbStairs https://leetcode.com/problems/climbing-stairs/
// Ref: https://leetcode.com/problems/climbing-stairs/solutions/25299/basically-it-s-a-fibonacci/
/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
func ClimbStairs(n int) int {
	return climbStairsMemo(n)
}

func climbStairsMemo(n int) int {
	memo := make(map[int]int)

	var climbFunc func(stair int) int
	climbFunc = func(stair int) int {
		// Base Case: Reach the start of the stair
		if stair == 1 || stair == 2 {
			return stair
		}

		if val, ok := memo[stair]; ok {
			return val
		}

		memo[stair] = climbFunc(stair-1) + climbFunc(stair-2)
		return memo[stair]
	}

	return climbFunc(n)
}

func climbStairsTable(n int) int {
	// It's necessary for initial check because of the edge case, like n = 1 which will lead to index out of range [2] with length 2
	if n == 1 || n == 2 {
		return n
	}

	// DP Table is always longer 1 of the target size
	table := make([]int, n+1)
	// Set 1 as standard default of table[0]
	table[0] = 1
	table[1] = 1
	table[2] = 2

	// The cap of the loop is "less than and equal to" rather than just less than
	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[n]
}

func climbStairsExercise(n int) int {
	return 0
}

/*
Time Complexity O(2^n) of Depth First Search (DFS) where "2" is derived from the ways to climb (2 decision tree) and n of the height of the tree
Time Complexity O(η^n) where η is Fibonacci constant := 1.83
Time Limit Exceeded!!!
*/
func climStairsRecursivelyImpl(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return climStairsRecursivelyImpl(n-1) + climStairsRecursivelyImpl(n-2)
	}
}

// Too Complicated....
func climStairsIteratively(n int) int {
	// First, define Base Cases
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// For Simplest Case of N = 3
	waysToTwoStepBeforeN := 1
	waysToOneStepBeforeN := 2

	for i := 3; i < n; i++ {
		waysToNStairs := waysToTwoStepBeforeN + waysToOneStepBeforeN
		waysToTwoStepBeforeN = waysToOneStepBeforeN
		waysToOneStepBeforeN = waysToNStairs
	}

	return waysToTwoStepBeforeN + waysToOneStepBeforeN
}
