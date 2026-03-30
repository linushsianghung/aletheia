package pattern

import "github.com/linushung/aletheia/Concepts/Data_Structure/Graph"

/*
Write a function, undirectedPath, that take in an array of edges for an undirected graph and 2 nodes (nodeA, NodeB).
The function should return a boolean indicating whether there exist a path between nodeA and nodeB.

Related Problem: 1971. Find if Path Exists in Graph: https://leetcode.com/problems/find-if-path-exists-in-graph/description/
*/
func undirectedPath(edges [][]string, nodeA, nodeB string) bool {
	graph := Graph.BuildGraph(edges)
	return undirectedPathDFS(graph, nodeA, nodeB, make(map[string]bool))
}

func undirectedPathDFS(graph map[string][]string, src, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true

	for _, neighbor := range graph[src] {
		if undirectedPathDFS(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}

func undirectedPathExercise(graph map[string][]string, src, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}
	if visited[src] {
		return false
	}
	visited[src] = true

	for _, neighbor := range graph[src] {
		if undirectedPathDFS(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}
