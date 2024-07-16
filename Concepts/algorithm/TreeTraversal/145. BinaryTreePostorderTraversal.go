package TreeTraversal

import "github.com/linushung/aletheia/leetcode"

// PostorderTraversal https://leetcode.com/problems/binary-tree-postorder-traversal/description/
// Ref: https://www.youtube.com/watch?v=QhszUQhGGlA
/* Given the root of a binary tree, return the postorder traversal of its nodes' values. */
func PostorderTraversal(root *leetcode.TreeNode) []int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	return postorderTraversalRecursively(root)
	// return postorderTraversalIteratively(root)
}

func postorderTraversalRecursively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)

	var localRecursiveFunc func(root *leetcode.TreeNode)
	localRecursiveFunc = func(root *leetcode.TreeNode) {
		if root == nil {
			return
		}

		localRecursiveFunc(root.Left)
		localRecursiveFunc(root.Right)
		result = append(result, root.Val)
	}

	localRecursiveFunc(root)
	return result
}

func postorderTraversalIteratively(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	stack := []*leetcode.TreeNode{root}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Because of Post Order Traversal which root node will always be visited last,
		// so just processed the node now but add new nodes to the head of the result.
		result = append([]int{current.Val}, result...)

		// Based on the example, the Left subtree has to be put first
		// Stack: [Left, Right,...] => Result: [Left, Right, Root]
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return result
}
