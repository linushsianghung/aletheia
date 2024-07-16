package testdome

import "fmt"

/**
Implement the uniqueNames function. When passed two slices of names, it will return a slice containing the names that
appear in either or both slices. The returned slice should have no duplicates.

For example, calling uniqueNames([]string{"Ava", "Emma", "Olivia"}, []string{"Olivia", "Sophia", "Emma"}) should return
a slice containing Ava, Emma, Olivia, and Sophia in any order.
 */
func uniqueNames(a, b []string) []string {
	aggreMap := make(map[string]struct{})
	result := make([]string, 0, len(aggreMap))

	comb := append(a, b...)
	for _, s := range comb {
		aggreMap[s] = struct{}{}
	}

	for k := range aggreMap {
		result = append(result, k)
	}

	return result
}

func MergeNames() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
