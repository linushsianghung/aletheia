package Tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
/*
You are given the root of a binary search tree (BST) and an integer val.
Find the node in the BST that the node's value equals val and return the subtree rooted with that node.
If such a node does not exist, return null.
*/
func searchBST(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	return searchBSIteratively(root, val)
}

func searchBSTRecursively(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBSTRecursively(root.Left, val)
	} else {
		return searchBSTRecursively(root.Right, val)
	}
}

func searchBSTRecursivelyExercise(root *leetcode.TreeNode, val int) *leetcode.TreeNode {

	return nil
}

func searchBSIteratively(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}

func searchBSIterativelyExercise(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	return nil
}

// Just for clarifying the concept of tree search
func searchTreeRecursively(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	// First search the node from left children nodes
	node := searchTreeRecursively(root.Left, val)
	if node != nil {
		return node
	}
	return searchTreeRecursively(root.Right, val)
}
