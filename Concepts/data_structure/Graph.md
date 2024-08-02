# Graph
Ref:
- [Graph Data Structure with 3 Javascript Implementations](https://www.youtube.com/watch?v=e4RezPkq3UI)
- [Graph Theory Tutorial from a Google Engineer](https://www.youtube.com/watch?v=09_LlHjoEiY)

## Definition
Graph is a data structure made up of **nodes**/**vertices** and **edges** or the connections between nodes. Tree is a special kind of graph which has only 1 root and 1 unique path between any 2 nodes

## Graph Representation & Implementation
```Golang
type Vertice struct { Val string }
type Edge struct {From, To string}

vertices := []Vertice{
    Vertice {"A"},
    Vertice {"B"},
    Vertice {"C"},
    Vertice {"D"},
    Vertice {"E"},
}

edges := [
    Edge{"A", "B"},
    Edge{"A", "D"}
    Edge{"B", "C"}
    Edge{"C", "D"}
    Edge{"C", "E"}
    Edge{"D", "E"}
]

func findAdjacentVertices(v Vertice) {

}

func isConnected(v1, v2 Vertice) {

}

```

