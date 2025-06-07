package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
*/
func kthSmallest(root *leetcode.TreeNode, k int) int {
	return kthSmallestExercise(root, k)
}

func kthSmallestIteratively(root *leetcode.TreeNode, k int) int {
	result, stack := make([]int, 0), make([]*leetcode.TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Val)
			if len(result) == k {
				return result[k-1]
			}

			current = current.Right
		}
	}

	return 0
}

func kthSmallestRecursively(root *leetcode.TreeNode, k int) int {
	result := make([]int, 0)

	var inorderFunc func(node *leetcode.TreeNode)
	inorderFunc = func(node *leetcode.TreeNode) {
		if node == nil || len(result) == k {
			return
		}

		inorderFunc(node.Left)
		result = append(result, node.Val)
		inorderFunc(node.Right)
	}
	inorderFunc(root)

	return result[k-1]
}

func kthSmallestExercise(root *leetcode.TreeNode, k int) int {

	return 0
}
