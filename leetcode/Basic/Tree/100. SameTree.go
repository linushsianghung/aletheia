package Tree

import "github.com/linushung/aletheia/leetcode"

// IsSameTree https://leetcode.com/problems/same-tree/description/
// Ref: https://www.youtube.com/watch?v=vRbbcKXCxOw
/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/
func IsSameTree(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {
	// Define Base Cases
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func isSameTreeExercise(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {

	return false
}
