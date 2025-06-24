package Tree

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/LeetCode75/BinaryTree_DFS"
)

// LowestCommonAncestor https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree
/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/
func LowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return LowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return LowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

// Related Problem: 236. Lowest Common Ancestor of a Binary Tree: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/?envId=leetcode-75
func lowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	return BinaryTree_DFS.LowestCommonAncestor(root, p, q)
}
