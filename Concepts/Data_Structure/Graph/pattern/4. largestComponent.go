package pattern

/*
Write a function, largestComponent, that takes in the adjacency list of an undirected graph.
The function should return the size of the largest connected component in the graph.

Largest Component: file://./pics/LargestComponent.png
*/

// largestComponent finds the size of the largest connected component in a graph.
func largestComponent(graph map[int][]int) int {
	maxCount := 0
	visited := make(map[int]bool)

	// Iterate through each node of the graph as a potential starting point.
	for node := range graph {
		// exploreSize will find the size of the component containing 'node'.
		// If 'node' has already been visited, its component has already been counted, and exploreSize will return 0.
		count := exploreSize(graph, node, visited)
		maxCount = max(maxCount, count)
	}

	return maxCount
}

// exploreSize uses Depth-First Search (DFS) to find the number of nodes in a single connected component starting from a given node.
func exploreSize(graph map[int][]int, node int, visited map[int]bool) int {
	// Base case: If the node has already been visited, we don't count it again.
	if visited[node] {
		return 0
	}
	// Mark the node as visited to prevent infinite loops and recounting.
	visited[node] = true

	count := 1 // Start the count for this component at 1 (for the current node).
	// Recursively visit all neighbors and add their sizes to the count.
	for _, neighbor := range graph[node] {
		count += exploreSize(graph, neighbor, visited)
	}

	return count
}

func largestComponentExercise(graph map[int][]int) int {
	return 0
}

func exploreSizeExercise(graph map[int][]int, node int, visited map[int]bool) int {
	return 0
}
