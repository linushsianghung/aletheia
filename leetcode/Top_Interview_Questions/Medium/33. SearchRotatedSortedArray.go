package Medium

import "github.com/linushung/aletheia/Concepts/Algorithm/Binary_Search"

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func search(nums []int, target int) int {
	return Binary_Search.SearchRotate(nums, target)
}
