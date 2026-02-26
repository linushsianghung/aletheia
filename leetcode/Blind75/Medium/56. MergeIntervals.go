package Medium

import "github.com/linushung/aletheia/Concepts/Algorithm/Merging"

// https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	return Merging.Merge(intervals)
}
