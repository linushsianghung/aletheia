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
	note := make([]int, len(cost))
	for i := range note {
		note[i] = -1
	}

	var localFunc func(index int) int
	localFunc = func(index int) int {
		if index < 0 {
			return 0
		}
		if index == 0 || index == 1 {
			return cost[index]
		}

		if value := note[index]; value > -1 {
			return value
		}

		note[index] = cost[index] + min(localFunc(index-1), localFunc(index-2))
		return note[index]
	}

	return min(localFunc(len(cost)-1), localFunc(len(cost)-2))
}

func minCostClimbingStairsStraight(cost []int) int {
	note := make([]int, len(cost))
	for i := range note {
		note[i] = -1
	}

	var localFunc func(index int) int
	localFunc = func(index int) int {
		if index >= len(cost) {
			return 0
		}
		if value := note[index]; value > -1 {
			return value
		}

		// Recursively traverse to the end first, then return current cost with whichever smaller one to get the minimum cost over all
		note[index] = cost[index] + min(localFunc(index+1), localFunc(index+2))
		return note[index]
	}

	// As statement, we can either start from the step with index 0, or the step with index 1.
	return min(localFunc(0), localFunc(1))
}
