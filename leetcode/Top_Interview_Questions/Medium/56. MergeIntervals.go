package Medium

import "sort"

// https://leetcode.com/problems/merge-intervals/
/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/
func merge(intervals [][]int) [][]int {
	// Step 1. Sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Step 2. Create a Stack for processing
	resultStack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		// Step 3.  1 element each from both Stack (for earlier interval) & Intervals (for later interval) for merging
		// intervalA is a earlier interval because it's pop from Stack
		intervalA := resultStack[len(resultStack)-1]
		resultStack = resultStack[:len(resultStack)-1]
		// intervalA is a later interval because it's polled from Queue
		intervalB := intervals[0]
		intervals = intervals[1:]

		// Step 4. Process both elements based on the scenario
		if intervalA[1] < intervalB[0] {
			// A.end < B.start => Non-Overlapping
			resultStack = append(resultStack, intervalA)
			resultStack = append(resultStack, intervalB)
		} else {
			// A.end >= B.start => Create a new interval for merging
			newInterval := []int{intervalA[0], max(intervalA[1], intervalB[1])}
			resultStack = append(resultStack, newInterval)
		}
	}

	return resultStack
}

func mergeExercise(intervals [][]int) [][]int {
	return nil
}
