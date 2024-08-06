package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
*/
func kthSmallest(root *leetcode.TreeNode, k int) int {
	result, stack := make([]int, 0), []*leetcode.TreeNode{root}

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Val)
			if len(result) == k {
				return result[k-1]
			}
			root = current.Right
		}
	}

	return 0
}
