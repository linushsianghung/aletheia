package Medium

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/insert-interval/
/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	stack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		intervalA := stack[len(stack)-1]
		stack := stack[:len(stack)-1]

		intervalB := intervals[0]
		intervals = intervals[1:]

		if intervalA[1] < intervalB[0] {
			stack = append(stack, intervalA)
			stack = append(stack, intervalB)
		} else {
			intervalC := []int{intervalA[0], max(intervalA[1], intervalB[1])}
			stack = append(stack, intervalC)
		}

	}

	fmt.Printf("Final Stack: %+v\n", stack)
	return stack
}
