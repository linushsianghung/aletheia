package Heap

import "container/heap"

// https://leetcode.com/problems/kth-largest-element-in-an-array/
/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?
*/
func findKthLargest(nums []int, k int) int {
	myHeap := &IntHeap{}
	heap.Init(myHeap)

	for _, num := range nums {
		heap.Push(myHeap, num)

		if myHeap.Len() > k {
			heap.Pop(myHeap)
		}
	}

	return heap.Pop(myHeap).(int)
}
