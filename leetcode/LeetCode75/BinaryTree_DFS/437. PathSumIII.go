package BinaryTree_DFS

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/path-sum-iii/description/?envId=leetcode-75
/*
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).
*/
func pathSum(root *leetcode.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// The search could start from any node in the tree with sum, we should see every node as root and do a dfs. Thus the $O(n^2)$ time complexity
	return pathSumDP(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func pathSumDP(node *leetcode.TreeNode, remain int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val == remain {
		count++
	}

	return count + pathSumDP(node.Left, remain-node.Val) + pathSumDP(node.Right, remain-node.Val)
}

func pathSumDPExercise(node *leetcode.TreeNode, remain int) int {
	return 0
}
