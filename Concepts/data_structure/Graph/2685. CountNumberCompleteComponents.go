package Graph

// https://leetcode.com/problems/count-the-number-of-complete-components/description/
/*
You are given an integer n. There is an undirected graph with n vertices, numbered from 0 to n - 1.
You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting vertices ai and bi.
Return the number of complete connected components of the graph.

- A connected component is a subgraph of a graph in which there exists a path between any two vertices, and no vertex of the subgraph shares an edge with a vertex outside the subgraph.
- A connected component is said to be complete if there exists an edge between every pair of its vertices.
*/
func countCompleteComponents(n int, edges [][]int) int {
	count := 0
	graph := BuildGraph(edges)

	var nestedHelper func(node int, note map[int]bool) bool
	nestedHelper = func(node int, note map[int]bool) bool {
		if note[node] {
			return false
		}
		note[node] = true

		for _, neighbor := range graph[node] {
			nestedHelper(neighbor, note)
		}

		return true
	}

	note := make(map[int]bool)
	for i := range n {
		if nestedHelper(i, note) {
			count++
		}
	}

	return count
}
