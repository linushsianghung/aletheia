package BinarySearch

// https://leetcode.com/problems/find-peak-element/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/find-peak-element/solutions/1290642/intuition-behind-conditions-complete-explanation-diagram-binary-search/
/*
A peak element is an element that is strictly greater than its neighbors.
Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.

Analysis:
For a mid-element, there could be three possible cases :
Case 1 : mid-element is equal to the peak element ( Observation : mid-element is greater than its neighbors )
Case 2 : mid-element lies on the right of our result peak ( Observation : Our peak element search space is left side )
Case 3 : mid-element lies on the left. ( Observation : Our peak element search space is right side )

For example:
4 is in the middle of _ _ _ 5 4 6 _ _ _. So, whichever side we move in this case, we are guaranteed to find a peak.
Let's take the left side first, the number previous to 5 can either be smaller or greater than 5. If it's smaller, 5 is a peak.
Or if it's larger, similarly look at the number previous to it. If the numbers keep on increasing on the left side, i.e. 8 7 6 5 4 6 _ _ _, the left-most would be a peak.

PS:
Binary search sorted assumption is necessary only when there is only 1 "unique value" in the array. In this case, there could be multiple, and we don't care which one to return.
*/
func findPeakElement(nums []int) int {
	// Check single element
	if len(nums) == 1 {
		return 0
	}

	// Base cases:
	// The array could be strictly increasing or decreasing, and it's given that nums[-1] = nums[n] = -∞.
	// Also, it's allowed to return any of the possible peaks, we could add a condition to check whether the 1st element/last element could be the peak.
	if nums[0] > nums[1] {
		return 0
	}

	n := len(nums)
	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	left, right := 1, n-2
	for left <= right {
		mid := left + (right-left)/2
		// Case 1
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
			// Case 2
		} else if nums[mid] < nums[mid-1] {
			right = mid - 1
			// Case 3
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		}
	}

	return -1
}
