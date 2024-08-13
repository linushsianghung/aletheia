package pattern

/*
Write a function, hasPath, that take in an object representing the adjacent list of a directed acyclic graph and 2 nodes (src, dst).
The function should return a boolean indicating whether there exist a directed path between the src and dst nodes.

Has Path: [file://./pics/HasPath.png]

Analysis:
- Time complexity: O(n)
- Space Complexity: O(e)
where n & v is the number of nodes and edges
*/
func hasPathRecursively(graph map[string][]string, src, dst string) bool {
	if src == dst {
		return true
	}

	for _, neighbor := range graph[src] {
		if hasPathRecursively(graph, neighbor, dst) {
			return true
		}
	}

	return false
}

func hasPathBFS(graph map[string][]string, src, dst string) bool {
	queue := []string{src}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == dst {
			return true
		}

		queue = append(queue, graph[current]...)
	}

	return false
}
