package pattern

import "github.com/linushung/aletheia/Concepts/data_structure/Graph"

/*
Write a function, shortestPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB).
The function should return the length of the shortest path between A and B. Consider the length as the number of edges in the path, not the number of nodes.
If there is no path between A and B, then return -1. You can assume that A and B exist as nodes in the graph.
*/
func shortestPath(edge [][]string, nodeA, nodeB string) int {
	graph := Graph.BuildGraph(edge)
	level, queue, note := 0, []string{nodeA}, make(map[string]bool)

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == nodeB {
				return level
			}

			note[current] = true
			for _, node := range graph[current] {
				if !note[node] {
					queue = append(queue, node)
				}
			}
		}

		level++
	}

	return level
}
