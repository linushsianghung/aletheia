package BinaryTree_DFS

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/?envId=leetcode-75
/*
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself).”
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	// if p or q is exactly the whole tree's root, return the root (because root is the lowest node).
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// There are 4 scenarios:
	if left == nil && right == nil {
		return nil
	} else if left != nil && right != nil {
		// If left != nil && right != nil, it means p & q locate on different side and the lca is root itself
		return root
	} else {
		// If both p & q locate on 1 side (either left == nil || right == nil), the lca is p || q which ever is not nil
		if left == nil {
			return right
		}
		return left
	}
}

func lowestCommonAncestorExercise(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	return nil
}
