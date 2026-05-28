package Merging

import (
	"sort"
)

// Merge https://leetcode.com/problems/merge-intervals
/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/
func Merge(intervals [][]int) [][]int {
	// Step 1. Sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Step 2. Create a Stack for processing
	stack := [][]int{intervals[0]}
	intervals = intervals[1:]

	for len(intervals) > 0 {
		// Step 3.  1 element each from both Stack (for earlier interval) & Intervals (for later interval) for merging
		// intervalA is a earlier interval because it's pop from Stack
		intervalA := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// intervalA is a later interval because it's polled from Queue
		intervalB := intervals[0]
		intervals = intervals[1:]

		// Step 4. Process both elements based on the scenario
		if intervalA[1] < intervalB[0] {
			// A.end < B.start => Non-Overlapping
			stack = append(stack, intervalA)
			stack = append(stack, intervalB)
		} else {
			// A.end >= B.start => Create a new interval for merging
			newInterval := []int{intervalA[0], max(intervalA[1], intervalB[1])}
			stack = append(stack, newInterval)
		}
	}

	return stack
}

func mergeExercise(intervals [][]int) [][]int {
	return nil
}

/*
Analysis:
1.  Concept Introduction/Explanation:
    The "Merge Intervals" problem asks us to consolidate a collection of intervals that overlap. For example, if we have `[1,3]` and `[2,6]`, they overlap and can be merged into `[1,6]`. If we then have `[8,10]`, it doesn't overlap with `[1,6]`, so it remains separate.

    The core idea behind solving this problem efficiently is a **greedy approach** combined with **sorting**.

    Imagine you have a timeline, and you're marking segments on it.
    *   If you sort all your segments by their starting points, you ensure that when you look at a segment, any segment that *could* overlap with it from the left has already been considered.
    *   Once sorted, you can iterate through the intervals and maintain a list of merged intervals.

    Here's the step-by-step process for the standard greedy approach:
    1.  **Sort the intervals:** The most crucial step is to sort all given intervals based on their `start` times. If two intervals have the same `start` time, their `end` times can be used as a secondary sorting criterion, but for this problem, sorting by `start` time alone is sufficient. This ensures that we process intervals from left to right on our conceptual timeline.
    2.  **Initialize a result list:** Create an empty list (e.g., `ArrayList` in Java, `[]int` slice in Go) to store the non-overlapping merged intervals.
    3.  **Add the first interval:** Add the first interval from the *sorted* input list to your result list. This interval will be the first candidate for merging.
    4.  **Iterate and merge:** Loop through the remaining intervals (starting from the second one) in the sorted input list. For each `current` interval:
        *   **Get the `lastMerged` interval:** This is the last interval that was added to your result list.
        *   **Check for overlap:** Compare `current.start` with `lastMerged.end`.
            *   **If `current.start <= lastMerged.end`:** An overlap exists! This means the `current` interval starts before or at the same time the `lastMerged` interval ends. To merge them, we update the `lastMerged.end` to be the maximum of `lastMerged.end` and `current.end`. The `lastMerged.start` remains unchanged.
            *   **If `current.start > lastMerged.end`:** No overlap. The `current` interval starts after the `lastMerged` interval ends. This means `current` is a new, distinct interval that cannot be merged with the previous ones. So, add `current` to your result list as a new entry.
    5.  **Return the result:** Once all intervals have been processed, the result list will contain all the non-overlapping merged intervals.

    **Analogy:** Imagine you're drawing lines on a piece of paper. You sort all the lines by their starting points. You draw the first line. Then, for every subsequent line, you check if it starts before your current drawn line ends. If it does, you just extend your current drawn line to cover the new line's end point (if it goes further). If it starts *after* your current drawn line ends, you pick up your pen and start drawing a completely new line.

2.  Complexity Analysis:
    *   **Time Complexity:**
        *   **Sorting:** The dominant factor is sorting the `N` intervals. Using efficient sorting algorithms (like `Arrays.sort` in Java or `sort.Slice` in Go, which are typically Timsort or Quicksort-based) takes `O(N log N)` time.
        *   **Merging:** After sorting, we iterate through the intervals once. Each interval is processed (compared and potentially merged or added) in `O(1)` time. This iteration takes `O(N)` time.
        *   **Total Time Complexity:** `O(N log N) + O(N) = O(N log N)`.
    *   **Space Complexity:**
        *   **Sorting:** The space complexity for sorting depends on the specific algorithm used. In Java, `Arrays.sort` for objects (which `int[][]` effectively is) can use `O(N)` space for temporary storage (e.g., for merge sort). In Go, `sort.Slice` can also use `O(log N)` to `O(N)` depending on the implementation.
        *   **Result List:** In the worst-case scenario (e.g., no intervals overlap), we might store all `N` intervals in our result list. This requires `O(N)` space.
        *   **Total Space Complexity:** `O(N)`.

3.  Interview Suggestion:
    *   **Start with the brute force (and why it's bad):** Briefly mention that checking every pair of intervals would be `O(N^2)`, which is too slow. This shows you consider alternatives.
    *   **Key Insight - Sorting:** Immediately pivot to the idea of sorting. Explain *why* sorting by start time is crucial (it simplifies the overlap check to only the last merged interval).
    *   **Greedy Strategy:** Clearly articulate the greedy approach: "Once sorted, we can iterate through the intervals, maintaining a list of merged intervals. For each new interval, we try to merge it with the *last* interval in our merged list."
    *   **Walk through an example:** Use a small example (e.g., `[[1,3],[8,10],[2,6],[15,18]]`) to demonstrate the sorting and merging steps.
    *   **Edge Cases:** Discuss empty input, single interval, all intervals overlapping, no intervals overlapping.
    *   **Complexity Analysis:** Be prepared to derive and explain the `O(N log N)` time and `O(N)` space complexities.
    *   **Code Structure:** In Java, using `ArrayList<int[]>` is idiomatic for dynamic resizing, then converting to `int[][]` at the end. In Go, using `[][]int` and `append` works similarly.
    *   **Review of current Go implementation:** Your current Go implementation uses a "stack" in a slightly unconventional way by popping from the stack and polling from the remaining `intervals` slice. While it achieves the correct result, the standard greedy approach (as described above and implemented in the Java code below) is generally simpler to reason about and implement. The standard approach typically involves initializing the result with the first interval and then iterating from the second, comparing with the *last* element of the result list.

Implementation Method (Java):
```java
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

// This class structure is typical for LeetCode solutions in Java
class Solution {
    public int[][] merge(int[][] intervals) {
        // Handle edge case: if the input array is null or empty, return an empty array.
        if (intervals == null || intervals.length == 0) {
            return new int[0][0];
        }

        // Step 1: Sort the intervals by their start times.
        // This is the most crucial step. Sorting ensures that we process intervals from left to right on the timeline, simplifying overlap checks.
        // Comparator.comparingInt(a -> a[0]) creates a comparator that compares the first element (start time) of each int array.
        Arrays.sort(intervals, Comparator.comparingInt(a -> a[0]));

        // Step 2: Initialize a list to store the merged intervals.
        // We use an ArrayList because the number of merged intervals is not known beforehand, and it allows dynamic resizing.
        List<int[]> mergedIntervals = new ArrayList<>();

        // Step 3: Add the first interval from the sorted list to our merged list.
        // This interval serves as our initial 'last merged interval' candidate.
        mergedIntervals.add(intervals[0]);

        // Step 4: Iterate through the rest of the sorted intervals, starting from the second one.
        for (int i = 1; i < intervals.length; i++) {
            // Get the last interval that was added to our 'mergedIntervals' list. This is the interval we will attempt to merge with the 'current' interval.
            int[] lastMerged = mergedIntervals.get(mergedIntervals.size() - 1);

            // Get the current interval from the input array that we are processing.
            int[] current = intervals[i];

            // Check for overlap:
            // An overlap exists if the start time of the 'current' interval is less than or equal to the end time of the 'lastMerged' interval.
            if (current[0] <= lastMerged[1]) {
                // Overlap detected: Merge the 'current' interval into the 'lastMerged' interval.
                // We update the end time of 'lastMerged' to be the maximum of its current end and the end time of the 'current' interval. The start time of 'lastMerged' remains unchanged as it's already the earliest start.
                lastMerged[1] = Math.max(lastMerged[1], current[1]);
            } else {
                // No overlap: The 'current' interval does not overlap with the 'lastMerged' interval.
                // This means 'current' is a new, distinct interval that should be added to our list.
                mergedIntervals.add(current);
            }
        }

        // Step 5: Convert the List of int arrays back to a 2D int array and return.
        // The toArray method with a pre-sized array argument is efficient for this conversion.
        return mergedIntervals.toArray(new int[mergedIntervals.size()][]);
    }
}
```
*/
