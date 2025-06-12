package Graphs_DFS

// https://leetcode.com/problems/evaluate-division/description/?envId=leetcode-75
/*
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and
values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.
You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.
Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.
Note: The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Build the Presentation of Graph.
	// Ex: equations = [["a","b"],["b","c"]], values = [2.0,3.0]
	// {
	//		a: { b: 2 },
	// 		b: { a: 1/2, c: 3 }
	//		c: { b: 1/3 }
	// }
	graph := make(map[string]map[string]float64)
	for i, equation := range equations {
		if _, ok := graph[equation[0]]; !ok {
			graph[equation[0]] = make(map[string]float64)
		}
		if _, ok := graph[equation[1]]; !ok {
			graph[equation[1]] = make(map[string]float64)
		}
		graph[equation[0]][equation[1]] = values[i]
		graph[equation[1]][equation[0]] = 1 / values[i]
	}

	result := make([]float64, len(queries))
	// Looping each query
	for i, query := range queries {
		// Below order of validation is really really matter!!!
		if _, ok := graph[query[0]]; !ok {
			result[i] = -1
			continue
		}
		// Special Cases. It couldn't be the first validation which will be failed at query like [x","x"]
		if query[0] == query[1] {
			result[i] = 1
			continue
		}

		// Based on the left of query as start, looping through each equation of map
		// Ex: b: { a: 1/2, c: 3 }
		for next, value := range graph[query[0]] {
			visited := make(map[string]bool)
			visited[query[0]] = true

			stop, value := explore(graph, query[1], next, value, visited)
			result[i] = value
			// If stop is true, which means find out the answer of this query, just break out of the for loop to finish the exploration
			if stop {
				break
			}
		}
	}

	return result
}

func explore(graph map[string]map[string]float64, destination, current string, product float64, visited map[string]bool) (bool, float64) {
	// Base case. Return true if it finds the answer
	if current == destination {
		return true, product
	}

	// Using DFS concept to find the answer recursively
	for next, value := range graph[current] {
		// Using a note to prevent from loop infinitely, like a: { b: 2 } and b: { a: 1/2, c: 3 }
		if ok := visited[next]; ok {
			continue
		}
		visited[next] = true

		stop, result := explore(graph, destination, next, product*value, visited)
		// If stop is true, which manes find out the answer, just return the result; otherwise explore next equation
		if stop {
			return true, result
		}
	}

	return false, -1
}
