package Tree

import (
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/univalued-binary-tree/
/*
A binary tree is uni-valued if every node in the tree has the same value.
Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.

Analysis:
Because we need to check the value of one node (say, root node) with each other child node, it's really unnecessary
to use any kinds of traversal. Just check each node 1 by 1, left to right.
*/
func isUnivalTree(root *leetcode.TreeNode) bool {
	/* Basic init check for LinkedList and Tree, but it's not necessary because of the Constraints of this problem */
	if root == nil {
		return true
	}

	return isUnivalTreeSequentialOrder(root)
}

func isUnivalTreeSequentialOrder(root *leetcode.TreeNode) bool {
	// Check each node on the left side of the tree
	left := root.Left == nil || (root.Left.Val == root.Val && isUnivalTreeSequentialOrder(root.Left))
	// Then check each node on the right side of the tree
	right := root.Right == nil || (root.Right.Val == root.Val && isUnivalTreeSequentialOrder(root.Right))

	return left && right
}

func isUnivalTreeSequentialOrderExercise(root *leetcode.TreeNode) bool {
	return false
}

func isUnivalTreeLevelOrder(root *leetcode.TreeNode) bool {
	queue := []*leetcode.TreeNode{root}
	target := root.Val

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				if current.Left.Val != target {
					return false
				}
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				if current.Right.Val != target {
					return false
				}
				queue = append(queue, current.Right)
			}
		}
	}

	return true
}
