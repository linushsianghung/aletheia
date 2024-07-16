package Medium

import (
	"math"

	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/validate-binary-search-tree
// Ref: https://www.youtube.com/watch?v=yEwSGhSsT0U
/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.
*/
func isValidBST(root *leetcode.TreeNode) bool {
	// return isValidRecursively(root, math.MinInt, math.MaxInt)
	return isValidInOrderTraversal(root)
}

func isValidRecursively(root *leetcode.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	return root.Val > min && root.Val < max && isValidRecursively(root.Left, min, root.Val) && isValidRecursively(root.Right, root.Val, max)
}

func isValidInOrderTraversal(root *leetcode.TreeNode) bool {
	if root == nil {
		return true
	}

	prev := math.MinInt
	stack := make([]*leetcode.TreeNode, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			if node.Val <= prev {
				return false
			}
			prev = node.Val
			stack = stack[:len(stack)-1]
			root = node.Right
		}
	}

	return true
}

// Too expensive
func isSubTreeGreater(root *leetcode.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	return root.Val > val && isSubTreeGreater(root.Left, val) && isSubTreeGreater(root.Right, val)
}

// Too expensive
func isSubTreeLess(root *leetcode.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	return root.Val < val && isSubTreeLess(root.Left, val) && isSubTreeLess(root.Right, val)
}
