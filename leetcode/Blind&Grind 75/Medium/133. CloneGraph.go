package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/clone-graph/description/
/*
Given a reference of a node in a connected undirected graph. Return a deep clone (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.
class Node {
    public int val;
    public List<Node> neighbors;
}

Test case format:
- For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.
- An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
- The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

Analysis:
Time Complexity: O(V + E)
- V = Vertices (Nodes), E = Edges.
- Visiting every node once and iterating over every edge once. This is optimal.

Space Complexity: O(V)
- The visited map stores V nodes.
- The recursion stack can go up to O(V) in the worst case (a line graph).
*/
func cloneGraph(node *leetcode.Node) *leetcode.Node {
	return nil
}

// cloneGraphDFS implements the DFS (Depth First Search) approach which might risk a stack overflow on extremely deep graphs
func cloneGraphDFS(node *leetcode.Node) *leetcode.Node {
	// Map to store the mapping from Original Node -> Cloned Node
	visited := make(map[*leetcode.Node]*leetcode.Node)

	// DFS (Depth First Search) approach and use dfs as function name to follow standard graph terminology.
	var dfs func(node *leetcode.Node) *leetcode.Node
	dfs = func(node *leetcode.Node) *leetcode.Node {
		if node == nil {
			return nil
		}

		if n, ok := visited[node]; ok {
			return n
		}

		clone := &leetcode.Node{Val: node.Val}
		visited[node] = clone
		for _, child := range node.Children {
			clone.Children = append(clone.Children, dfs(child))
		}
		return clone
	}

	return dfs(node)
}

// cloneGraphBFS implements the BFS (Breadth First Search) approach.
// This uses a Queue instead of recursion.
func cloneGraphBFS(node *leetcode.Node) *leetcode.Node {
	if node == nil {
		return nil
	}

	// Map to store the mapping from Original Node -> Cloned Node
	visited := make(map[*leetcode.Node]*leetcode.Node)

	// Initialize the first node
	visited[node] = &leetcode.Node{Val: node.Val}
	queue := []*leetcode.Node{node}

	for len(queue) > 0 {
		// Pop from queue
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range current.Children {
			// If neighbor hasn't been cloned yet, clone it and add to queue
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &leetcode.Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			// Link the current clone to the neighbor clone
			visited[current].Children = append(visited[current].Children, visited[neighbor])
		}
	}

	return visited[node]
}

func cloneGraphExercise(node *leetcode.Node) *leetcode.Node {
	return nil
}
