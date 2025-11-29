# Sliding Window
Ref:
- [Sliding window technique](https://www.youtube.com/watch?v=p-ss2JNynmw)
- [Cracking Common Interview Algorithm Patterns: Sliding Window](https://www.youtube.com/watch?v=BM0mhAlvyQc)
- [NeetCode: Sliding Window](https://www.youtube.com/playlist?list=PLot-Xpze53leOBgcVsJBEGrHPd_7x_koV)

## Concepts:
- Sliding Window is an extension of the two-pointer approach where we use two pointers (left and right) to create a **window**. The problem will ask us to return the maximum or minimum sub-range that satisfies a given condition.
- The main idea behind the sliding window technique is to convert two nested loops into a single loop. Usually, the technique helps us to reduce the time complexity from O(n²) or O(n³) to O(n).This is done by maintaining a sliding window, which is a sub-array of the original array that is of a fixed size. The algorithm then iterates over the original array, updating the sliding window as it goes. This allows the algorithm to keep track of a contiguous sequence of elements in the original array, without having to iterate over the entire array multiple times.

## Sliding Window Problems:
### Fixed Size Sliding Window
This is generally the simpler of the two patterns.
   
**How to Identify**
The problem will explicitly give you the size of the window, often denoted as k. The question will be about finding something within a subarray or substring of that exact size.
- **Keywords**: `subarray of size k`, `substring of length k`, and `in any k elements`.
- **Example Problem**: Given an array of integers, find the maximum sum of any contiguous subarray of size k.

### The Approach (Template)
The core idea is to create an initial window, calculate the first result, and then slide the window one element at a time, making incremental updates.
- **Build the first window**: Process the first k elements of the array/string. Calculate your initial value (e.g., the sum, the character counts, etc.).
- **Slide and Update**: Loop from the k-th element to the end of the array. In each step:
  - **Add** the new element (the one at the right end of the new window).
  - **Remove** the old element (the one falling off the left end of the window).
  - **Update** your global answer (e.g., `max_sum = max(max_sum, current_sum)`).

### Dynamic Size Sliding Window
This pattern is more versatile and common in medium-to-hard problems.

### How to Identify
The problem will ask for the longest or shortest subarray/substring that satisfies a certain condition. The size of the window is not given to you; you have to find it.
- **Keywords**: `longest substring`, `shortest subarray`, `minimum length`, and `maximum size`.
- **Example Problem**: Given a string, find the length of the longest substring with no more than k distinct characters.

### The Approach (Template)
The core idea is to use two pointers (`winStart` and `winEnd`) to define the window. You expand the window by moving `winEnd` and shrink it by moving `winStart` whenever a condition is violated.
- **Initialize**: Set up your pointers (winStart = 0, winEnd = 0), a data structure to track the window's state (like a hash map), and a variable for your answer (e.g., maxLength = 0).
- **Expand**: Start a loop to move the right pointer from the beginning to the end of the array/string.
  - In each step, add the `winEnd` element to the window and update its state (e.g., increment its count in the map).
- **Shrink**: After expanding, check if the window is now **invalid** based on the problem's condition (e.g., `distinct_characters > k`).
  - If it's invalid, start a nested loop to move the `winStart` pointer to the right until the window becomes valid again.
  - While shrinking, remove the winStart element from the window and update its state (e.g., decrement its count in the map).
- **Update Answer**: After each valid expansion (and before a potential shrink), update the answer of candidate window (e.g., `maxLength = max(maxLength, winEnd - winStart + 1)`).

### Summary
| Feature      | Fixed Size Window                                              | Dynamic Size Window                                            |
|:-------------|:---------------------------------------------------------------|:---------------------------------------------------------------|
| Goal         | Find best value (max/min/avg) in a subarray of a given size k. | Find the longest/shortest subarray that satisfies a condition. |
| Window Size  | Constant, given in the problem.                                | Variable, changes based on a condition.                        |
| Movement     | Slides one step at a time.                                     | Expands with right pointer, shrinks with left pointer.         |
| Key Question | "What's the best result for a window of size k?"               | "What's the best window length that meets this criteria?"      |