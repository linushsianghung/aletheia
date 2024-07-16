package Tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/merge-two-binary-trees/
/*
You are given two binary trees root1 and root2.
Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.
You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new
value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.
Return the merged tree.

Note: The merging process must start from the root nodes of both trees.
*/
func mergeTrees(root1 *leetcode.TreeNode, root2 *leetcode.TreeNode) *leetcode.TreeNode {
	/* Basic init check for LinkedList and Tree */
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	/* Standard step for Tree algorithm: traverse Tree recursively */
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func mergeTreesExercise(root1 *leetcode.TreeNode, root2 *leetcode.TreeNode) *leetcode.TreeNode {

	return nil
}
