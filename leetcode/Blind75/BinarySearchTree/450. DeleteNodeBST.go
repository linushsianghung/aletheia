package BinarySearchTree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/delete-node-in-a-bst/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/delete-node-in-a-bst/solutions/821420/python-o-h-solution-explained
/*
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
Basically, the deletion can be divided into two stages:

- Search for a node to remove.
- If the node is found, delete the node.
*/
func deleteNode(root *leetcode.TreeNode, key int) *leetcode.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		// If key value is less than root value, go to left subtree to find the node
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		// If key value is greater than root value, go to right subtree to find the node
		root.Right = deleteNode(root.Right, key)
	} else {
		// Case 1: Node do not have any children then we just delete it and nothing else to do here.

		// Case 2: Node has Right child but do not have Left child, just skip this node and link its children with its parent
		if root.Left == nil {
			return root.Right
		}
		// Case 23: Node has Left child but do not have right child, just skip this node and link its children with its parent
		if root.Right == nil {
			return root.Left
		}

		if root.Right != nil && root.Left != nil {
			// Find the minimum node which is the most left node in right subtree
			temp := root.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			// Replace the value of current node with minimum value
			root.Val = temp.Val
			// Recursively go deeper to delete the node from which just get the minimum value
			root.Right = deleteNode(root.Right, root.Val)
		}
	}

	return root
}

func deleteNodeExercise(root *leetcode.TreeNode, key int) *leetcode.TreeNode {
	return root
}
