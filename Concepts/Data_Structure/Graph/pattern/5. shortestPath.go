package pattern

import "github.com/linushung/aletheia/Concepts/Data_Structure/Graph"

/*
Write a function, shortestPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB).
The function should return the length of the shortest path between A and B. Consider the length as the number of edges in the path, not the number of nodes.
If there is no path between A and B, then return -1. You can assume that A and B exist as nodes in the graph.
*/
// shortestPath finds the length of the shortest path between two nodes in a graph.
// It uses a Breadth-First Search (BFS) approach, which is ideal for finding the
// shortest path in an unweighted graph.
func shortestPath(edge [][]string, nodeA, nodeB string) int {
	graph := Graph.BuildGraph(edge)

	queue := []string{nodeA}
	// The visited map keeps track of nodes we've already added to the queue to avoid cycles and redundant processing.
	visited := make(map[string]bool)
	visited[nodeA] = true

	// level represents the distance (number of edges) from nodeA.
	level := 0
	for len(queue) > 0 {
		// Process all nodes at the current level.
		levelSize := len(queue)

		for range levelSize {
			current := queue[0]
			queue = queue[1:]

			if current == nodeB {
				return level
			}

			visited[current] = true
			// Add all unvisited neighbors to the queue for the next level.
			for _, neighbor := range graph[current] {
				if visited[neighbor] {
					continue
				}

				queue = append(queue, neighbor)
			}
		}
		level++
	}

	// If the queue becomes empty, and we haven't found nodeB, there is no path.
	return -1
}

func shortestPathExercise(edge [][]string, nodeA, nodeB string) int {
	return 0
}
