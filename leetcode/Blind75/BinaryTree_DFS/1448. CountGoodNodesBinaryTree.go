package BinaryTree_DFS

import (
	"github.com/linushung/aletheia/leetcode"
	"math"
)

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/?envId=leetcode-75
/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *leetcode.TreeNode) int {
	return goodNodesHelper(root)
}

func goodNodesHelper(root *leetcode.TreeNode) int {
	result := make([]int, 0)
	minVal := math.MinInt

	// Pre-Order Traversal
	var localRecursiveFunc func(root *leetcode.TreeNode, minVal int)
	localRecursiveFunc = func(root *leetcode.TreeNode, minVal int) {
		if root == nil {
			return
		}

		if root.Val >= minVal {
			result = append(result, root.Val)
			minVal = root.Val
		}
		localRecursiveFunc(root.Left, minVal)
		localRecursiveFunc(root.Right, minVal)
	}
	localRecursiveFunc(root, minVal)

	return len(result)
}
