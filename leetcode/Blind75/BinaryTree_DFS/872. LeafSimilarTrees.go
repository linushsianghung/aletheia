package BinaryTree_DFS

import (
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/leaf-similar-trees/?envId=leetcode-75
/*
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *leetcode.TreeNode, root2 *leetcode.TreeNode) bool {
	result1 := leafSimilarHelper(root1)
	result2 := leafSimilarHelper(root2)

	if len(result1) != len(result2) {
		return false
	}

	for i, num := range result1 {
		if num != result2[i] {
			return false
		}
	}

	return true
}

func leafSimilarHelper(root *leetcode.TreeNode) []int {
	result := make([]int, 0)

	// Post-Order Traversal
	var localRecursiveFunc func(root *leetcode.TreeNode) int
	localRecursiveFunc = func(root *leetcode.TreeNode) int {
		if root == nil {
			return -1
		}

		left := localRecursiveFunc(root.Left)
		right := localRecursiveFunc(root.Right)
		if left == -1 && right == -1 {
			result = append(result, root.Val)
		}
		return 0
	}

	localRecursiveFunc(root)
	return result
}
