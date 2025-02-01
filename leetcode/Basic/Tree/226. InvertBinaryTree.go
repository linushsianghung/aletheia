package Tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/invert-binary-tree/
// Ref: https://www.youtube.com/watch?v=OnSn2XEQ4MY (Depth First Search)
/* Given the root of a binary tree, invert the tree, and return its root. */
func invertTree(root *leetcode.TreeNode) *leetcode.TreeNode {
	invertTreeLevelOrderTraversal(root)
	return root
}

func invertTreeLevelOrderTraversal(root *leetcode.TreeNode) {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return
	}

	// Swap the left and right subtrees of root
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTreeLevelOrderTraversal(root.Left)
	invertTreeLevelOrderTraversal(root.Right)
}

func invertTreeLevelOrderTraversalExercise(root *leetcode.TreeNode) *leetcode.TreeNode {

	return nil
}
