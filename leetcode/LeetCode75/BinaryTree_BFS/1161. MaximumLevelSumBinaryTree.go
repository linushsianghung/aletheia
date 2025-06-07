package BinaryTree_BFS

import (
	"github.com/linushung/aletheia/leetcode"
	"math"
)

// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/description/?envId=leetcode-75
/*
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at level x is maximal.
*/
func maxLevelSum(root *leetcode.TreeNode) int {
	level, minLevel, maxSum := 0, 0, math.MinInt32
	queue := []*leetcode.TreeNode{root}

	for len(queue) > 0 {
		level++
		size, sum := len(queue), 0

		for range size {
			current := queue[0]
			queue = queue[1:]

			sum += current.Val
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			minLevel = level
		}
	}

	return minLevel
}

func maxLevelSumExercise(root *leetcode.TreeNode) int {
	return 0
}
