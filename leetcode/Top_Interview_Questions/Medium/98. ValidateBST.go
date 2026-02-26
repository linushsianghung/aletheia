package Medium

import (
	"math"

	"github.com/linushung/aletheia/leetcode"
)

// IsValidBST https://leetcode.com/problems/validate-binary-search-tree
// Ref: https://www.youtube.com/watch?v=yEwSGhSsT0U
/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

Analysis:
Time Complexity: O(N)
1. Traversal vs. Search:
	- Search ($O(\log N)$): When searching, you discard half the tree at every step (e.g., "target is smaller, so ignore the right side"). You only visit the nodes along a single path.
	- Validation ($O(N)$): To prove a tree is a valid BST, you must verify every single node. A tiny node deep in the bottom-right corner could be smaller than the root, which would invalidate the entire tree. Since you cannot skip any node, you must visit all $N$ nodes.
2. Worst Case: The worst case for validation is actually a valid BST (or one where the error is the very last node visited), because the algorithm has to traverse the entire structure to confirm it is correct.

*/
func IsValidBST(root *leetcode.TreeNode) bool {
	return isValidRecursively(root)
	//return isValidInOrderTraversal(root)
}

// This implementation will be failed if the tree contains math.MinInt or math.MaxInt as valid values (because val <= min returns false).
func isValidRecursively(node *leetcode.TreeNode) bool {
	var validate func(node *leetcode.TreeNode, min, max int) bool
	validate = func(node *leetcode.TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val >= max || node.Val <= min {
			return false
		}

		return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
	}

	return validate(node, math.MinInt, math.MaxInt)
}

// isValidBSTImprovement uses pointers to handle edge cases where node.Val equals math.MinInt/MaxInt.
func isValidBSTImprovement(root *leetcode.TreeNode) bool {
	var validate func(node *leetcode.TreeNode, min, max *int) bool
	validate = func(node *leetcode.TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}
		if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
			return false
		}
		return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
	}
	return validate(root, nil, nil)
}

func isValidBSTExercise(root *leetcode.TreeNode) bool {
	return false
}

// Time: O(N), Space: O(H) for stack.
func isValidInOrderTraversal(root *leetcode.TreeNode) bool {
	if root == nil {
		return true
	}

	var prev *int // optimizes space by not storing the full result slice
	//result := make([]int, 0)

	current := root
	stack := make([]*leetcode.TreeNode, 0)

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if prev != nil && current.Val <= *prev {
				return false
			}
			val := current.Val
			prev = &val
			current = current.Right
		}
	}

	return true
}

// Too expensive
func isSubTreeGreater(root *leetcode.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	return root.Val > val && isSubTreeGreater(root.Left, val) && isSubTreeGreater(root.Right, val)
}

// Too expensive
func isSubTreeLess(root *leetcode.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	return root.Val < val && isSubTreeLess(root.Left, val) && isSubTreeLess(root.Right, val)
}
