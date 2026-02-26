package Graph

// CountCompleteComponents https://leetcode.com/problems/count-the-number-of-complete-components/description/
/*
You are given an integer n. There is an undirected graph with n vertices, numbered from 0 to n - 1.
You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting vertices ai and bi.
Return the number of complete connected components of the graph.

- A connected component is a subgraph of a graph in which there exists a path between any two vertices, and no vertex of the subgraph shares an edge with a vertex outside the subgraph.
- A connected component is said to be complete if there exists an edge between every pair of its vertices.
*/
func CountCompleteComponents(n int, edges [][]int) int {
	// Build the graph using an adjacency list.
	// Since nodes are numbered 0 to n-1, a slice of slices is efficient.
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	count := 0
	visited := make([]bool, n)

	for i := range n {
		// If the node has already been visited, it belongs to a component we've already processed.
		if visited[i] {
			continue
		}

		// Start a traversal (BFS) to find all nodes in this component.
		// We need to track the number of nodes and the number of edges (sum of degrees) in this component.
		nodeCount := 0
		edgeCount := 0 // This will store the sum of degrees of all nodes in the component.

		queue := []int{i}
		visited[i] = true

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			nodeCount++
			// Add the degree of the current node. Note: This counts each edge twice (once for each end).
			edgeCount += len(graph[curr])

			for _, neighbor := range graph[curr] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}

		// A component is complete if every pair of distinct vertices is connected by an edge.
		// The number of edges in a complete graph with V vertices is V * (V - 1) / 2.
		// Since our edgeCount is the sum of degrees (2 * number of edges), we check if edgeCount == nodeCount * (nodeCount - 1).
		if edgeCount == nodeCount*(nodeCount-1) {
			count++
		}
	}

	return count
}

func countCompleteComponentsExercise(n int, edges [][]int) int {
	return 0
}
