package TreeTraversal

import "github.com/linushung/aletheia/leetcode"

// ZigzagLevelOrder https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right,
then right to left for the next level and alternate between).
*/
func ZigzagLevelOrder(root *leetcode.TreeNode) [][]int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	return zigzagLevelOrderIteratively(root)
	// return zigzagLevelOrderRecursively(root)
}

func zigzagLevelOrderIteratively(root *leetcode.TreeNode) [][]int {
	result := make([][]int, 0)
	queue := []*leetcode.TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		nodes := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				nodes[i] = current.Val
			} else {
				// Put value from the end of the slice
				nodes[len(nodes)-1-i] = current.Val
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		result = append(result, nodes)
		level++
	}

	return result
}

func zigzagLevelOrderRecursively(root *leetcode.TreeNode) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(root *leetcode.TreeNode, level int)
	localRecursiveFunc = func(root *leetcode.TreeNode, level int) {
		if root == nil {
			return
		}

		if len(result) == level {
			result = append(result, make([]int, 0))
		}

		if level%2 == 0 {
			result[level] = append(result[level], root.Val)
		} else {
			result[level] = append([]int{root.Val}, result[level]...)
		}
		localRecursiveFunc(root.Left, level+1)
		localRecursiveFunc(root.Right, level+1)
	}

	localRecursiveFunc(root, 0)
	return result
}
