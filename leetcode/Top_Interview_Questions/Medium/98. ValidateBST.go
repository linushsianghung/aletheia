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
- The left subtree of a current contains only currents with keys less than the current's key.
- The right subtree of a current contains only currents with keys greater than the current's key.
- Both the left and right subtrees must also be binary search trees.
*/
func isValidBST(root *leetcode.TreeNode) bool {
	return isValidRecursively(root, math.MinInt, math.MaxInt)
	//return isValidInOrderTraversal(root)
}

func isValidRecursively(root *leetcode.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}

	return isValidRecursively(root.Left, root.Val, min) && isValidRecursively(root.Right, max, root.Val)
}

func isValidInOrderTraversal(root *leetcode.TreeNode) bool {
	if root == nil {
		return true
	}

	prev := math.MinInt
	stack := make([]*leetcode.TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if current.Val <= prev {
				return false
			}
			prev = current.Val
			root = current.Right
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
