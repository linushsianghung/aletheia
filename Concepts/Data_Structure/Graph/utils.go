package Graph

/*
BuildGraph => Transferring Edge list to Adjacent list, which is easier for graph traversal
String:
- edges := [ ["i", "j"], ["k", "i"], ["m", "k"], ["k", "l"], ["o", "n"] ]
- graph := map[i:[j k] j:[i] k:[i m l] l:[k] m:[k] n:[o] o:[n]]
*/
func BuildGraph[T string | int](edges [][]T) map[T][]T {
	graph := make(map[T][]T)

	for _, edge := range edges {
		left, right := edge[0], edge[1]

		graph[left] = append(graph[left], right)
		graph[right] = append(graph[right], left)
	}

	return graph
}

func BuildGraphExercise[T string | int](edges [][]T) map[T][]T {
	return nil
}
