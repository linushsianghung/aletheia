package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/diameter-of-binary-tree
/*
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/
func diameterOfBinaryTree(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}

	maxLength := maxDepth(root.Left) + maxDepth(root.Right)
	lMaxLength := diameterOfBinaryTree(root.Left)
	rMaxLength := diameterOfBinaryTree(root.Right)

	return max(maxLength, max(lMaxLength, rMaxLength))
}

func maxDepth(node *leetcode.TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(maxDepth(node.Left), maxDepth(node.Right))
}
