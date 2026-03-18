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
	result1 := leafSimilarHelperIteratively(root1)
	result2 := leafSimilarHelperIteratively(root2)

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

func leafSimilarHelperRecursively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)

	var nestedFunc func(node *leetcode.TreeNode)
	nestedFunc = func(node *leetcode.TreeNode) {
		if node == nil {
			return
		}

		nestedFunc(node.Left)
		nestedFunc(node.Right)

		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
		}
	}

	nestedFunc(root)
	return result
}

func leafSimilarHelperIteratively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	stack := []*leetcode.TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Left == nil && current.Right == nil {
			result = append(result, current.Val)
			continue
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return result
}

func leafSimilarExercise(root1 *leetcode.TreeNode, root2 *leetcode.TreeNode) bool {
	return false
}

func leafSimilarHelperExercise(root *leetcode.TreeNode) []int {
	return nil
}
