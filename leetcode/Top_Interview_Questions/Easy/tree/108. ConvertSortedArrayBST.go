package tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Ref:
// - https://www.youtube.com/watch?v=0K0uCMYq5ng
// - https://www.youtube.com/watch?v=12omz-VAyRk
/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
*/
func sortedArrayToBST(nums []int) *leetcode.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return constructSubBST(nums, 0, len(nums)-1)
}

func constructSubBST(nums []int, left, right int) *leetcode.TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &leetcode.TreeNode{Val: nums[mid]}
	node.Left = constructSubBST(nums, left, mid-1)
	node.Right = constructSubBST(nums, mid+1, right)

	return node
}

func constructSubBSTExercise(nums []int, left, right int) *leetcode.TreeNode {
	return nil
}
