package Tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/path-sum
/*
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.
*/
func hasPathSum(root *leetcode.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Based on the description, this condition is how to define the root-to-leaf path with the sum equal to target
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func hasPathSumExercise(root *leetcode.TreeNode, targetSum int) bool {
	return false
}
