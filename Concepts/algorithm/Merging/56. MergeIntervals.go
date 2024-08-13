package Merging

import (
	"github.com/linushung/aletheia/leetcode/Top_Interview_Questions"
	"sort"
)

// https://leetcode.com/problems/merge-intervals
/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	resultStack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		intervalA := resultStack[len(resultStack)-1]
		resultStack = resultStack[:len(resultStack)-1]
		intervalB := intervals[0]
		intervals = intervals[1:]

		// Scenario1: Non-Overlapping
		if intervalA[1] < intervalB[0] {
			resultStack = append(resultStack, intervalA)
			resultStack = append(resultStack, intervalB)
		} else {
			newInterval := []int{intervalA[0], Top_Interview_Questions.Max(intervalA[1], intervalB[1])}
			resultStack = append(resultStack, newInterval)
		}
	}

	return resultStack
}
