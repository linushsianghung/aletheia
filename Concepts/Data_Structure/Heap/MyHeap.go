package Heap

// IntHeap implements heap.Interface and holds ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop need pointer receivers because they modify the slice's length, not just its contents.
func (h *IntHeap) Push(x any) {
	// Push is a pointer receiver as it modifies the slice's length.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	root := old[len(old)-1]
	*h = old[:len(old)-1]
	return root
}
