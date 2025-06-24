package Medium

import "github.com/linushung/aletheia/Concepts/Algorithm/Binary Search"

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	return Binary_Search.SearchRotate(nums, target)
}
