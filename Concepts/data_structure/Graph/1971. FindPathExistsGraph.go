package Graph

// https://leetcode.com/problems/find-if-path-exists-in-graph/description/
/*
There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive).
The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi.
Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.

You want to determine if there is a valid path that exists from vertex source to vertex destination.

Given edges and the integers n, source, and destination, return true if there is a valid path from source to destination, or false otherwise.
*/
func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := BuildGraph(edges)

	var validPathHelper func(src int, note map[int]bool) bool
	validPathHelper = func(src int, note map[int]bool) bool {
		if src == destination {
			return true
		}

		if note[src] {
			return false
		}
		note[src] = true

		for _, node := range graph[src] {
			if validPathHelper(node, note) {
				return true
			}
		}

		return false
	}

	return validPathHelper(source, make(map[int]bool))
}
