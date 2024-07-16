package Array

// https://leetcode.com/problems/merge-sorted-array/description/
/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n,
where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	tail1 := m - 1
	tail2 := n - 1
	tail := m + n - 1

	for ; tail1 >= 0 && tail2 >= 0; tail-- {
		if nums1[tail1] > nums2[tail2] {
			nums1[tail] = nums1[tail1]
			tail1--
		} else {
			nums1[tail] = nums2[tail2]
			tail2--
		}
	}

	for tail2 >= 0 {
		nums1[tail] = nums2[tail2]
		tail2--
		tail--
	}
}

func mergeExercise(nums1 []int, m int, nums2 []int, n int) {

}
