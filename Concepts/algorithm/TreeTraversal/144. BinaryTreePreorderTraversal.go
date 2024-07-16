package TreeTraversal

import "github.com/linushung/aletheia/leetcode"

// PreorderTraversal https://leetcode.com/problems/binary-tree-preorder-traversal/description/
// Ref: https://www.youtube.com/watch?v=afTpieEZXck
/* Given the root of a binary tree, return the preorder traversal of its nodes' values. */
func PreorderTraversal(root *leetcode.TreeNode) []int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	return preorderTraversalRecursively(root)
	// return preorderTraversalIteratively(root)
}

func preorderTraversalRecursively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)

	var localRecursiveFunc func(root *leetcode.TreeNode)
	localRecursiveFunc = func(root *leetcode.TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		localRecursiveFunc(root.Left)
		localRecursiveFunc(root.Right)
	}

	localRecursiveFunc(root)
	return result
}

func preorderTraversalIteratively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	stack := []*leetcode.TreeNode{root}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Pre-Order Traversal Visiting
		result = append(result, current.Val)

		// Put nodes to stack if the children nodes are not null
		// Based on the example, the Right subtree has to be put first
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

func preorderTraversalIterativelyAlt(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*leetcode.TreeNode, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			// Because it's Pre-Order Traversal, just visits the current node without bothering to store it
			result = append(result, root.Val)
			// Only store the right node to stack and move reference directly to the next level of the left node
			stack = append(stack, root.Right)
			root = root.Left
		} else {
			// Pop up the node from the tail of stack
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return result
}
