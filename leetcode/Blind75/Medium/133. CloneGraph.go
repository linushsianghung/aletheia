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
*/
func cloneGraph(node *leetcode.Node) *leetcode.Node {
	visited := make(map[*leetcode.Node]*leetcode.Node)

	var cloneFunc func(node *leetcode.Node) *leetcode.Node
	cloneFunc = func(node *leetcode.Node) *leetcode.Node {
		if node == nil {
			return nil
		}

		if n, ok := visited[node]; ok {
			return n
		}

		clone := &leetcode.Node{Val: node.Val}
		visited[node] = clone
		for _, child := range node.Children {
			clone.Children = append(clone.Children, cloneFunc(child))
		}
		return clone
	}

	return cloneFunc(node)
}
