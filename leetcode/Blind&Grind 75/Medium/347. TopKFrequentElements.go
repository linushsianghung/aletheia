package Medium

import "github.com/linushung/aletheia/Concepts/Data_Structure/Heap"

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequentElements(nums []int, k int) []int {
	return Heap.TopKFrequentElements(nums, k)
}
