package Tree

import (
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/univalued-binary-tree/
/*
A binary tree is uni-valued if every node in the tree has the same value.
Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.

Analysis:
Because we need to check the value of one node (say, root node) with each other child node, it's really unnecessary
to use any kinds of traversal. Just check each node 1 by 1, left to right.
*/
func isUnivalTree(root *leetcode.TreeNode) bool {
	return isUnivalTreeSequentialOrder(root)
}

func isUnivalTreeSequentialOrder(root *leetcode.TreeNode) bool {
	// Check each node on the left side of the tree
	left := root.Left == nil || (root.Left.Val == root.Val && isUnivalTreeSequentialOrder(root.Left))
	// Then check each node on the right side of the tree
	right := root.Right == nil || (root.Right.Val == root.Val && isUnivalTreeSequentialOrder(root.Right))

	return left && right
}

func isUnivalTreeSequentialOrderExercise(root *leetcode.TreeNode) bool {

	return false
}
