package BinaryTree_DFS

import "github.com/linushung/aletheia/leetcode"

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

	for i := range result1 {
		if result1[i] != result2[i] {
			return false
		}
	}

	return true
}

func leafSimilarHelper(root *leetcode.TreeNode) []int {
	result := make([]int, 0)

	// Post-Order Traversal
	var nestedFunc func(node *leetcode.TreeNode) int
	nestedFunc = func(node *leetcode.TreeNode) int {
		if node == nil {
			return -1
		}

		left := nestedFunc(node.Left)
		right := nestedFunc(node.Right)
		if left == -1 && right == -1 {
			result = append(result, node.Val)
		}
		return 0
	}

	nestedFunc(root)
	return result
}

func leafSimilarExercise(root1 *leetcode.TreeNode, root2 *leetcode.TreeNode) bool {
	return false
}

func leafSimilarHelperExercise(root *leetcode.TreeNode) []int {
	return nil
}
