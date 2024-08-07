# Graph
Ref:
- [Graph Data Structure with 3 Javascript Implementations](https://www.youtube.com/watch?v=e4RezPkq3UI)
- [Graph Theory Tutorial from a Google Engineer](https://www.youtube.com/watch?v=09_LlHjoEiY)

## Definition
Graph is a data structure made up of **nodes**/**vertices** and **edges** or the connections between nodes. Tree is a special kind of graph which has only 1 root and 1 unique path between any 2 nodes

## Graph Representation & Implementation
### Vertex with Edge List
```Golang
type Vertice struct { Val string }
type Edge struct {From, To string}

var (
	vertices = []Vertex{
		{"A"},
		{"B"},
		{"C"},
		{"D"},
		{"E"},
	}

	edges = []Edge{
		{"A", "B"},
		{"A", "D"},
		{"B", "C"},
		{"C", "D"},
		{"C", "E"},
		{"D", "E"},
	}
)

func findAdjacentVertices(v Vertice) []string {
    result := make([]string, 0)

    for _, edge := range edges {
        if edge.From == v.Val {
            result = append(result, edge.To)
        } else if edge.To == v. val {
            result = append(result, edge.From)
        }
    }

    return result
}

func isConnected(v1, v2 Vertice) bool {

}

```

### Adjacent Matrix
> A 2-D array filled out with 1's and 0's where each array represents a node and each index in the subarray represents a potential connection to another node
> The value at adjacentMatrix[node1][node2] indicates where there is a connection between node1 and node2
Analysis:
- Time Complexity to find adjacent nodes: O(v) where v is the number of vertices
- Time complexity to check if 2 nodes are connected: O(1)
- Space Complexity: O(v^2) where v is the number of vertices
=> **Great for there are many connections**
```Golang
type Vertex struct { Val string }
type Edge struct {From, To string}

var (
	vertices = []Vertex{
		{"A"},
		{"B"},
		{"C"},
		{"D"},
		{"E"},
	}

	edges = []Edge{
		{"A", "B"},
		{"A", "D"},
		{"B", "C"},
		{"C", "D"},
		{"C", "E"},
		{"D", "E"},
	}

	vertexIndex    = vertexIndexMap()
	adjacentMatrix = initAdjacentMatrix()
)

func vertexIndexMap() map[string]int {

	indexMap := make(map[string]int)
	for i, v := range vertices {
		indexMap[v.Val] = i
	}

	// IndexMap: map[A:0 B:1 C:2 D:3 E:4]
	return indexMap
}

func initAdjacentMatrix() [][]int {
	indexMap := vertexIndexMap()

	matrix := make([][]int, len(vertices))
	for i, _ := range matrix {
		matrix[i] = make([]int, len(vertices))
	}

	for _, edge := range edges {
		matrix[indexMap[edge.To]][indexMap[edge.From]] = 1
		matrix[indexMap[edge.From]][indexMap[edge.To]] = 1
	}

	// Matrix: [[0 1 0 1 0] [1 0 1 0 0] [0 1 0 1 1] [1 0 1 0 1] [0 0 1 1 0]]
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

func isConnected(v1, v2 Vertice) bool {

}

```