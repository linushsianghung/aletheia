package Merging

import "sort"

// https://leetcode.com/problems/merge-intervals
/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/
func merge(intervals [][]int) [][]int {
	// Sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Using stack to merge sorted intervals 1 by 1
	stack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		// Pop interval from stack as intervalA
		intervalA := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Pop interval from slice as intervalB
		intervalB := intervals[0]
		intervals = intervals[1:]

		// Scenario: Non-Overlapping
		if intervalA[1] < intervalB[0] {
			stack = append(stack, intervalA)
			stack = append(stack, intervalB)
			continue
		}

		// Scenario: Overlapping => Create new intervalC for merging
		newInterval := []int{intervalA[0], max(intervalA[1], intervalB[1])}
		stack = append(stack, newInterval)
	}

	return stack
}

func mergeExercise(intervals [][]int) [][]int {
	return nil
}
