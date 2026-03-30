package pattern

import "github.com/linushung/aletheia/Concepts/Data_Structure/Graph"

/*
Write a function, connectedComponentsCount, that take in an adjacent list of an undirected graph.
The function should return the number of connected components within the graph

Connected Components Count: file://./pics/ConnectedComponentsCount.png
*/
func connectedComponentsCount(graph map[int][]int) int {
	count := 0
	visited := make(map[int]bool)

	for node := range graph {
		if explore(graph, node, visited) {
			count++
		}
	}

	return count
}

func explore(graph map[int][]int, node int, visited map[int]bool) bool {
	if visited[node] {
		return false
	}
	visited[node] = true

	for _, neighbor := range graph[node] {
		explore(graph, neighbor, visited)
	}

	return true
}

func connectedComponentsCountExercise(graph map[int][]int) int {
	return 0
}

func exploreExercise(graph map[int][]int, node int, visited map[int]bool) bool {
	return false
}

// Related Problem: 2685. Count the Number of Complete Components: https://leetcode.com/problems/count-the-number-of-complete-components/description/
func countCompleteComponents(n int, edges [][]int) int {
	return Graph.CountCompleteComponents(n, edges)
}
