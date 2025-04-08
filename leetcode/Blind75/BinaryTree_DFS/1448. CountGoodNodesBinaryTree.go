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
	return goodNodesRecursively(root)
}

func goodNodesRecursively(root *leetcode.TreeNode) int {
	count := 0

	var nestedFunc func(node *leetcode.TreeNode, minVal int)
	nestedFunc = func(node *leetcode.TreeNode, minVal int) {
		if node == nil {
			return
		}

		if node.Val >= minVal {
			count++
			minVal = node.Val
		}
		nestedFunc(node.Left, minVal)
		nestedFunc(node.Right, minVal)
	}
	nestedFunc(root, math.MinInt)

	return count
}

func goodNodesRecursivelyExercise(root *leetcode.TreeNode) int {
	return 0
}

func goodNodesPreOrderTemplate(root *leetcode.TreeNode) int {
	result := make([]int, 0)

	// Pre-Order Traversal
	var nestedFunc func(root *leetcode.TreeNode, minVal int)
	nestedFunc = func(root *leetcode.TreeNode, minVal int) {
		if root == nil {
			return
		}

		if root.Val >= minVal {
			result = append(result, root.Val)
			minVal = root.Val
		}
		nestedFunc(root.Left, minVal)
		nestedFunc(root.Right, minVal)
	}

	nestedFunc(root, math.MinInt)
	return len(result)
}
