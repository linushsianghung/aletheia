package Medium

import "github.com/linushung/aletheia/leetcode"

// KthSmallest https://leetcode.com/problems/kth-smallest-element-in-a-bst/
/* Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree. */
func KthSmallest(root *leetcode.TreeNode, k int) int {
	return kthSmallestExercise(root, k)
}

func kthSmallestIteratively(root *leetcode.TreeNode, k int) int {
	stack := make([]*leetcode.TreeNode, 0)
	count, current := 0, root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count++
			if count == k {
				return current.Val
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
