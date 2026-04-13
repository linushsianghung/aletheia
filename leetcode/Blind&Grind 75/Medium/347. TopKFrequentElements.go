package Medium

import (
	"github.com/linushung/aletheia/Concepts/Data_Structure/HashTable"
)

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequentElements(nums []int, k int) []int {
	return HashTable.TopKFrequentElements(nums, k)
}
