# Graph
Ref:
- [Graph Data Structure with 3 Javascript Implementations](https://www.youtube.com/watch?v=e4RezPkq3UI)
- [Graph Theory Tutorial from a Google Engineer](https://www.youtube.com/watch?v=09_LlHjoEiY)

## Definition
Graph is a data structure made up of **nodes**/**vertices** and **edges** or the connections between nodes.
Tree is a special kind of graph which has only 1 root and 1 unique path between any 2 nodes

## Graph Representation & Implementation
### Vertex with Edge List
> Analysis:
> - Time Complexity to find adjacent nodes: O(e) where e is the number of edges
> - Time Complexity to check if 2 nodes are connected: O(e) where e is the number of edges
> - Space Complexity: O(v + e) where v is the number of vertices and e is the number of edges

```Golang
package data_structure

import "slices"

type Vertex struct{ Val string }

var (
	vertices = []Vertex{
		{"A"},
		{"B"},
		{"C"},
		{"D"},
		{"E"},
	}

	edges = [][]string{
		{"A", "B"},
		{"A", "D"},
		{"B", "C"},
		{"C", "D"},
		{"C", "E"},
		{"D", "E"},
	}
)

func findAdjacentVertices(v Vertex) []string {
	result := make([]string, 0)

	for _, edge := range edges {
		if edge[0] == v.Val {
			result = append(result, edge[1])
		} else if edge[1] == v.Val {
			result = append(result, edge[0])
		}
	}

	return result
}

func isConnected(v1, v2 Vertex) bool {
	for _, edge := range edges {
		if slices.Contains(edge, v1.Val) && slices.Contains(edge, v2.Val) {
			return true
        }
	}

	return false
}
```

### Adjacent Matrix
> A 2-D array filled out with 1's and 0's where each array represents a node and each index in the subarray represents a potential connection to another node
> The value at adjacentMatrix[node1][node2] indicates where there is a connection between node1 and node2
> Analysis:
> - Time Complexity to find adjacent nodes: O(v) where v is the number of vertices
> - Time complexity to check if 2 nodes are connected: O(1)
> - Space Complexity: O(v^2) where v is the number of vertices
=> **Great for there are many connections**
```Golang
package data_structure

type Vertex struct { Val string }

var (
	vertices = []Vertex{
		{"A"},
		{"B"},
		{"C"},
		{"D"},
		{"E"},
	}

	edges = [][]string{
		{"A", "B"},
		{"A", "D"},
		{"B", "C"},
		{"C", "D"},
		{"C", "E"},
		{"D", "E"},
	}

	vertexIndex    = vertexIndexMap()       // IndexMap: map[A:0 B:1 C:2 D:3 E:4]
	adjacentMatrix = initAdjacentMatrix()	// Matrix: [[0 1 0 1 0] [1 0 1 0 0] [0 1 0 1 1] [1 0 1 0 1] [0 0 1 1 0]]
)

func vertexIndexMap() map[string]int {

	indexMap := make(map[string]int)
	for i, v := range vertices {
		indexMap[v.Val] = i
	}
	
	return indexMap
}

func initAdjacentMatrix() [][]int {
	indexMap := vertexIndexMap()

	matrix := make([][]int, len(vertices))
	for i := range matrix {
		matrix[i] = make([]int, len(vertices))
	}

	for _, edge := range edges {
		matrix[indexMap[edge[0]]][indexMap[edge[1]]] = 1
		matrix[indexMap[edge[1]]][indexMap[edge[0]]] = 1
	}
	
	return matrix
}

func findAdjacentVertices(v Vertex) []string {
	indexMap := vertexIndexMap()
	vIdx := indexMap[v.Val]
	result := make([]string, 0)

	for i, v := range adjacentMatrix[vIdx] {
		if v == 1 {
			result = append(result, vertices[i].Val)
		}
	}

	return result
}

func isConnected(v1, v2 Vertex) bool {
    return adjacentMatrix[vertexIndex[v1.Val]][vertexIndex[v2.Val]] == 1
}
```

### Adjacent List
> For every node, store a list of what nodes it's connected to
> Analysis:
> - Time Complexity to find adjacent nodes: O(1)
> - Time complexity to check if 2 nodes are connected: O(logv) if each adjacent row is sorted
> - Space Complexity: O(e) where v is the number of vertices