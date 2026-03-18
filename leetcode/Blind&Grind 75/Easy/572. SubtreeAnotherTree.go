package Easy

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Basic/Tree"
)

// https://leetcode.com/problems/subtree-of-another-tree/
/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.
*/
func isSubtree(root *leetcode.TreeNode, subRoot *leetcode.TreeNode) bool {
	if isSameSubTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameSubTree(node *leetcode.TreeNode, subNode *leetcode.TreeNode) bool {
	if node == nil || subNode == nil {
		return node == subNode
	}

	return node.Val == subNode.Val && isSameSubTree(node.Left, subNode.Left) && isSameSubTree(node.Right, subNode.Right)
}

func isSubtreeExercise(root *leetcode.TreeNode, subRoot *leetcode.TreeNode) bool {
	return false
}

func isSameSubTreeExercise(node *leetcode.TreeNode, subNode *leetcode.TreeNode) bool {
	return false
}

// Related Problem: 100. Same Tree: https://leetcode.com/problems/same-tree
func isSameTree100(p *leetcode.TreeNode, q *leetcode.TreeNode) bool { return Tree.IsSameTree(p, q) }
