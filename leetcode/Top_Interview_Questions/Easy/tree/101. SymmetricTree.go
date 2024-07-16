package tree

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Basic/Tree"
)

// https://leetcode.com/problems/symmetric-tree/
// Ref: https://www.youtube.com/watch?v=Mao9uzxwvmc
/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/
func isSymmetric(root *leetcode.TreeNode) bool {
	return LevelOrderCompare(root.Left, root.Right)
}

func LevelOrderCompare(left, right *leetcode.TreeNode) bool {
	// Define Base Cases
	if left == nil || right == nil {
		return left == right
	}

	return left.Val == right.Val && LevelOrderCompare(left.Left, right.Right) && LevelOrderCompare(left.Right, right.Left)
}

func LevelOrderCompareExercise(left, right *leetcode.TreeNode) bool {

	return false
}

// Related Topic: 100 - Same Tree: https://leetcode.com/problems/same-tree/
func isSameTree(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {
	return Tree.IsSameTree(p, q)
}
