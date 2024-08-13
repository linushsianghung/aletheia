package Graph

/*
edges := [ ["i", "j"], ["k", "i"], ["m", "k"], ["k", "l"], ["o", "n"] ]
graph := map[i:[j k] j:[i] k:[i m l] l:[k] m:[k] n:[o] o:[n]]
*/
func buildGraph(edges [][]string) map[string][]string {
	result := make(map[string][]string)

	for _, edge := range edges {
		if _, ok := result[edge[0]]; !ok {
			result[edge[0]] = make([]string, 0)
		}
		result[edge[0]] = append(result[edge[0]], edge[1])

		if _, ok := result[edge[1]]; !ok {
			result[edge[1]] = make([]string, 0)
		}
		result[edge[1]] = append(result[edge[1]], edge[0])
	}

	return result
}
