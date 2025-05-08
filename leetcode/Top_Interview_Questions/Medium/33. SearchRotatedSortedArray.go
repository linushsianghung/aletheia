package Medium

import "github.com/linushung/aletheia/Concepts/algorithm/BinarySearch"

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func search(nums []int, target int) int {
	return BinarySearch.SearchRotate(nums, target)
}
