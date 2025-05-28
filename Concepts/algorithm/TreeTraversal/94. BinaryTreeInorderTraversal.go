package TreeTraversal

import "github.com/linushung/aletheia/leetcode"

// InorderTraversal https://leetcode.com/problems/binary-tree-inorder-traversal/
// Ref: https://www.youtube.com/watch?v=g_S5WuasWUE
/* Given the root of a binary tree, return the postorder traversal of its nodes' values. */
func InorderTraversal(root *leetcode.TreeNode) []int {
	// result := make([]int, 0)
	// recursiveTraversal(root. &result)
	// return result

	return inorderTraversalRecursively(root)
}

func inorderTraversalRecursively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)

	var nestedFunc func(node *leetcode.TreeNode)
	nestedFunc = func(node *leetcode.TreeNode) {
		/* Basic init check for LinkedList and Tree */
		if node == nil {
			return
		}

		// Standard way for Tree Traversal recursively to the subtree
		nestedFunc(node.Left)
		result = append(result, node.Val)
		nestedFunc(node.Right)
	}
	nestedFunc(root)

	return result
}

func inorderTraversalRecursivelyExercise(root *leetcode.TreeNode) []int {
	return nil
}

func inorderTraversalIteratively(root *leetcode.TreeNode) []int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack, current := make([]*leetcode.TreeNode, 0), root

	for current != nil || len(stack) > 0 {
		// Continue go deeper into the left side of the tree until null
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Val)
			// Then go into the right side of the tree
			current = current.Right
		}
	}

	return result
}

func inorderTraversalIterativelyExercise(root *leetcode.TreeNode) []int {
	if root != nil {
		return nil
	}

	return nil
}

func recursiveTraversal(root *leetcode.TreeNode, result *[]int) {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return
	}

	recursiveTraversal(root.Left, result)
	*result = append(*result, root.Val)
	recursiveTraversal(root.Right, result)
}
