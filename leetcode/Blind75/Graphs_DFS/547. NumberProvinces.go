package Graphs_DFS

// https://leetcode.com/problems/number-of-provinces/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/number-of-provinces/solutions/3420785/number-of-provinces/?envType=study-plan-v2&envId=leetcode-75
/*
There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b,
and city b is connected directly with city c, then city a is connected indirectly with city c.
A province is a group of directly or indirectly connected cities and no other cities outside the group.
You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.
*/

func findCircleNum(isConnected [][]int) int {
	count, note := 0, make(map[int]bool)

	var nestedFunc func(city int, is map[int]bool) bool
	nestedFunc = func(city int, note map[int]bool) bool {
		if note[city] {
			return false
		}
		note[city] = true

		for i, connection := range isConnected[city] {
			if i != city && connection != 0 {
				nestedFunc(i, note)
			}
		}

		return true
	}

	for i := range isConnected {
		if nestedFunc(i, note) {
			count++
		}
	}

	return count
}
