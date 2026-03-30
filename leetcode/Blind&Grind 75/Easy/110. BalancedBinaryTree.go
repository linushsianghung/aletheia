package Easy

import (
	"math"

	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/balanced-binary-tree/description/
/* Given a binary tree, determine if it is height-balanced. */
func isBalanced(root *leetcode.TreeNode) bool {

	var balanceFunc func(node *leetcode.TreeNode) int
	balanceFunc = func(node *leetcode.TreeNode) int {
		if node == nil {
			return 0
		}

		left := balanceFunc(node.Left)
		right := balanceFunc(node.Right)

		if left == -1 || right == -1 {
			return -1
		}

		if math.Abs(float64(left-right)) > 1 {
			return -1
		}

		return 1 + max(left, right)
	}

	return balanceFunc(root) != -1
}

func isBalancedExercise(root *leetcode.TreeNode) bool {
	return false
}
