package pattern

/*
Write a function, connectedComponentsCount, that take in an adjacent list of an undirected graph.
The function should return the number of connected components within the graph

Connected Components Count: [file://./pics/ConnectedComponentsCount.png]
*/
func connectedComponentsCount(graph map[string][]string) int {
	sum := 0

	for node := range graph {
		if exploreCount(graph, node, make(map[string]bool)) {
			sum++
		}
	}
	return sum
}

func exploreCount(graph map[string][]string, node string, visited map[string]bool) bool {
	if _, ok := visited[node]; ok {
		return false
	}
	visited[node] = true

	for _, neighbor := range graph[node] {
		exploreCount(graph, neighbor, visited)
	}

	return true
}
