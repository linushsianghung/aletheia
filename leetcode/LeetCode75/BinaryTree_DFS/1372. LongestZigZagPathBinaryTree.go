package BinaryTree_DFS

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/solutions/534418/java-c-dfs-solution-with-comment-o-n-clean-code/
/*
You are given the root of a binary tree. A ZigZag path for a binary tree is defined as follows:

- Choose any node in the binary tree and a direction (right or left).
- If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
- Change the direction from right to left or from left to right.
- Repeat the second and third steps until you can't move in the tree.
- Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.
*/
func longestZigZag(root *leetcode.TreeNode) int {
	maxLen := 0

	var nestedFunc func(node *leetcode.TreeNode, path int, fromLeft bool)
	nestedFunc = func(node *leetcode.TreeNode, path int, fromLeft bool) {
		if node == nil {
			maxLen = max(maxLen, path)
			return
		}

		if fromLeft {
			nestedFunc(node.Right, path+1, false) // Continue to right zigzag path
			nestedFunc(node.Left, 0, true)        // Start a new zigzag path from left
		} else {
			nestedFunc(node.Left, path+1, true) // Continue to left zigzag path
			nestedFunc(node.Right, 0, false)    // Start a new zigzag path from right
		}
	}

	// In order to handle some special cases, like root = [1], using -1 as start point
	nestedFunc(root, -1, true)
	return maxLen
}

func longestZigZagExercise(root *leetcode.TreeNode) int {
	return 0
}
