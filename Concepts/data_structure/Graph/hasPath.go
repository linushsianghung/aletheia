
// Has Path: [file://./pics/HasPath.png]
/*
Analysis:
- Time complexity: O(n)
- Space Complexity: O(e)
where n & v is the number of nodes and edges
*/
package Graph

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
	return  false
}

