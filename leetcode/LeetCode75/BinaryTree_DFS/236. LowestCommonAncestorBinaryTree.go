package BinaryTree_DFS

import "github.com/linushung/aletheia/leetcode"

// LowestCommonAncestor https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/?envId=leetcode-75
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
func LowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	// if p or q is exactly the whole tree's root, return the root (because root is the lowest node).
	if root == p || root == q || root == nil {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
		// Cannot fine p || q in this branch, just return nil
		return nil
	} else if left != nil && right != nil {
		// If left != nil && right != nil, it means p & q locate on different side and the lca is current root
		return root
	} else {
		// If one side is nil, it means cannot find p || q on that side, just return the side whichever is not nil
		if left == nil {
			return right
		}
		return left
	}
}

func lowestCommonAncestorExercise(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	return nil
}
