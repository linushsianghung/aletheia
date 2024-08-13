package pattern

import "github.com/linushung/aletheia/Concepts/data_structure/Graph"

/*
Write a function, undirectedPath, that take in an array of edges for an undirected graph and 2 nodes (nodeA, NodeB).
The function should return a boolean indicating whether there exist a path between nodeA and nodeB.
*/
func undirectedPath(edges [][]string, nodeA, nodeB string) bool {
	graph := Graph.buildGraph(edges)
	return hasPath(graph, nodeA, nodeB, make(map[string]bool))
}

func hasPath(graph map[string][]string, src, dst string, visited map[string]bool) bool {
	if _, ok := visited[src]; ok {
		return false
	}
	visited[src] = true

	if src == dst {
		return true
	}

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}
