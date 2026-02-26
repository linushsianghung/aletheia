package pattern

import "github.com/linushung/aletheia/Concepts/Data_Structure/Graph"

/*
Write a function, hasPath, that take in an object representing the adjacent list of a directed acyclic graph and 2 nodes (src, dst).
The function should return a boolean indicating whether there exist a directed path between the src and dst nodes.

Has Path: file://./pics/HasPath.png

Analysis:
- Time complexity: O(n)
- Space Complexity: O(e)
where n & e is the number of nodes and edges
*/
func hasPath(graph map[string][]string, src, dst string) bool {
	return hasPathDFS(graph, src, dst)
}

func hasPathDFS(graph map[string][]string, src, dst string) bool {
	if src == dst {
		return true
	}

	for _, neighbor := range graph[src] {
		if hasPathDFS(graph, neighbor, dst) {
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

func hasPathExercise(graph map[string][]string, src, dst string) bool {
	return false
}

// Related Problem: 1971. Find if Path Exists in Graph: https://leetcode.com/problems/find-if-path-exists-in-graph/description/
func validPath(n int, edges [][]int, source int, destination int) bool {
	return Graph.ValidPath(n, edges, source, destination)
}
