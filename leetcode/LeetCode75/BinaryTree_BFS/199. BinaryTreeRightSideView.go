package BinaryTree_BFS

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/binary-tree-right-side-view/?envId=leetcode-75
/*
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
*/

func rightSideView(root *leetcode.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	queue := []*leetcode.TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := range size {
			current := queue[0]
			queue = queue[1:]

			if i == 0 {
				result = append(result, current.Val)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
		}
	}

	return result
}

func rightSideViewExercise(root *leetcode.TreeNode) []int {
	return nil
}
