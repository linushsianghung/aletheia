package Medium

import "sort"

// https://leetcode.com/problems/insert-interval/
/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	// return insertSimple(intervals, newInterval)
	return insertSmartCompare(intervals, newInterval)
}

func insertSmartCompare(intervals [][]int, newInterval []int) [][]int {
	// Use index to control the progress otherwise it requires extra code to handle empty intervals case when using "for range"
	index, result := 0, make([][]int, 0)

	// Just add each interval into result before meeting the newInterval
	for index < len(intervals) && intervals[index][1] < newInterval[0] {
		result = append(result, intervals[index])
		index++
	}

	// Merge overlapping intervals with newInterval
	for index < len(intervals) && intervals[index][0] <= newInterval[1] {
		newInterval = []int{
			min(intervals[index][0], newInterval[0]),
			max(intervals[index][1], newInterval[1]),
		}
		index++
	}
	// Add newly merging newInterval
	result = append(result, newInterval)

	// Add the rest of intervals
	result = append(result, intervals[index:]...)

	return result
}

func insertSmartCompareExercise(intervals [][]int, newInterval []int) [][]int {
	return nil
}

func insertSimple(intervals [][]int, newInterval []int) [][]int {
	// Simply add a new interval to intervals then it the same as basic Merge Intervals question
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	stack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		intervalA := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		intervalB := intervals[0]
		intervals = intervals[1:]

		if intervalA[1] < intervalB[0] {
			stack = append(stack, intervalA)
			stack = append(stack, intervalB)
			continue
		}

		intervalC := []int{intervalA[0], max(intervalA[1], intervalB[1])}
		stack = append(stack, intervalC)
	}

	return stack
}
