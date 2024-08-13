package pattern

/*
Write a function, largestComponent, that takes in the adjacency list of an undirected graph.
The function should return the size of the largest connected component in the graph.

Largest Component: [file://./pics/LargestComponent.png]
*/

func largestComponent(graph map[string][]string) int {
	maxSize := 0
	for node := range graph {
		size := exploreSize(graph, node, make(map[string]bool))
		if maxSize < size {
			maxSize = size
		}
	}

	return maxSize
}

func exploreSize(graph map[string][]string, node string, visited map[string]bool) int {
	if visited[node] {
		return 0
	}
	visited[node] = true

	size := 0
	for _, neighbor := range graph[node] {
		size += exploreSize(graph, neighbor, visited)
	}

	return size
}
