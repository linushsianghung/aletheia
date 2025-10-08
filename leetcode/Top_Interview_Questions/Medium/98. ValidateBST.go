package Medium

import (
	"math"

	"github.com/linushung/aletheia/leetcode"
)

// IsValidBST https://leetcode.com/problems/validate-binary-search-tree
// Ref: https://www.youtube.com/watch?v=yEwSGhSsT0U
/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.
*/
func IsValidBST(root *leetcode.TreeNode) bool {
	return isValidRecursively(root, math.MinInt, math.MaxInt)
	//return isValidInOrderTraversal(root)
}

func isValidRecursively(node *leetcode.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val >= max || node.Val <= min {
		return false
	}

	return isValidRecursively(node.Left, min, node.Val) && isValidRecursively(node.Right, node.Val, max)
}

func isValidRecursivelyExercise(node *leetcode.TreeNode, min, max int) bool {
	return false
}

func isValidInOrderTraversal(root *leetcode.TreeNode) bool {
	if root == nil {
		return true
	}

	current := root
	stack, result := make([]*leetcode.TreeNode, 0), make([]int, 0)

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(result) > 0 && result[len(result)-1] >= current.Val {
				return false
			}
			result = append(result, current.Val)
			current = current.Right
		}
	}

	return true
}

func isValidInOrderTraversalExercise(root *leetcode.TreeNode) bool {
	return false
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
