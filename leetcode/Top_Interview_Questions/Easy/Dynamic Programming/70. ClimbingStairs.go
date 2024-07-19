package Dynamic_Programming

// https://leetcode.com/problems/climbing-stairs/
// Ref: https://leetcode.com/problems/climbing-stairs/solutions/25299/basically-it-s-a-fibonacci/
/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
func climbStairs(n int) int {
	return climStairsMemoriseDP(n, make(map[int]int))
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

func climStairsMemoriseDP(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		memo[n] = climStairsMemoriseDP(n-1, memo) + climStairsMemoriseDP(n-2, memo)
		return memo[n]
	}
}

func climStairsBottomUpDP(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	tabu := make([]int, 0)
	tabu = append(tabu, 0)
	tabu = append(tabu, 1)
	tabu = append(tabu, 2)

	for i := 3; i <= n; i++ {
		tabu = append(tabu, tabu[i-1]+tabu[i-2])
	}

	return tabu[n]
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
