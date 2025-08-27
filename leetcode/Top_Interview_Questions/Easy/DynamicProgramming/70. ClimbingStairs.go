package DynamicProgramming

// ClimbStairs https://leetcode.com/problems/climbing-stairs/
// Ref: https://leetcode.com/problems/climbing-stairs/solutions/25299/basically-it-s-a-fibonacci/
/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
func ClimbStairs(n int) int {
	return climStairsMemoriseDP(n)
}

func climStairsMemoriseDP(n int) int {
	var climbFunc func(staircase int, note map[int]int) int
	climbFunc = func(staircase int, note map[int]int) int {
		if steps, ok := note[staircase]; ok {
			return steps
		}

		switch staircase {
		case 1:
			return 1
		case 2:
			return 2
		default:
			note[staircase] = climbFunc(staircase-1, note) + climbFunc(staircase-2, note)
			return note[staircase]
		}
	}

	return climbFunc(n, make(map[int]int))
}

func climStairsMemoriseDPExercise(n int, memo map[int]int) int {
	return 0
}

func climStairsBottomUpDP(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1
	table[2] = 2

	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[n]
}

func climStairsBottomUpDPExercise(n int) int {
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
