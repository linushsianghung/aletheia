package DP_1D

// https://leetcode.com/problems/min-cost-climbing-stairs/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/min-cost-climbing-stairs/solutions/476388/4-ways-step-by-step-from-recursion-top-down-dp-bottom-up-dp-fine-tuning/
/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.
*/
// DP Memorisation
func minCostClimbingStairs(cost []int) int {
	note := make(map[int]int)

	var costFunc func(index int) int
	costFunc = func(index int) int {
		if index < 0 {
			return 0
		}
		if val, ok := note[index]; ok {
			return val
		}

		note[index] = cost[index] + min(costFunc(index-1), costFunc(index-2))
		return note[index]
	}

	return min(costFunc(len(cost)-1), costFunc(len(cost)-2))
}

func minCostClimbingStairsExercise(cost []int) int {
	note := make([]int, len(cost))
	for i := range note {
		note[i] = -1
	}

	var minCostFunc func(index int) int
	minCostFunc = func(index int) int {
		if index < 0 {
			return 0
		}

		if note[index] != -1 {
			return note[index]
		}

		note[index] = cost[index] + min(minCostFunc(index-1), minCostFunc(index-2))
		return note[index]
	}

	return minCostFunc(len(cost) - 1)
}

func minCostClimbingStairsStraight(cost []int) int {
	note := make([]int, len(cost))
	for i := range note {
		note[i] = -1
	}

	var costFunc func(index int) int
	costFunc = func(index int) int {
		if index >= len(cost) {
			return 0
		}
		if value := note[index]; value > -1 {
			return value
		}

		// Recursively traverse to the end first, then return current cost with whichever smaller one to get the minimum cost over all
		note[index] = cost[index] + min(costFunc(index+1), costFunc(index+2))
		return note[index]
	}

	// As statement, we can either start from the step with index 0, or the step with index 1.
	return min(costFunc(0), costFunc(1))
}
