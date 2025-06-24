package Intervals

import "sort"

// EraseOverlapIntervals https://leetcode.com/problems/non-overlapping-intervals/description/?envId=leetcode-75
/*
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.
*/
func EraseOverlapIntervals(intervals [][]int) int {
	// Sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	stack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		intervalA := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		intervalB := intervals[0]
		intervals = intervals[1:]

		if intervalA[1] <= intervalB[0] {
			stack = append(stack, intervalA)
			stack = append(stack, intervalB)
		} else {
			count++
			intervalC := []int{intervalA[0], min(intervalA[1], intervalB[1])}
			stack = append(stack, intervalC)
		}
	}

	return count
}

func eraseOverlapIntervalsExercise(intervals [][]int) int {
	return 0
}
