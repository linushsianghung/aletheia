package BinarySearchTree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/delete-node-in-a-bst/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/delete-node-in-a-bst/solutions/821420/python-o-h-solution-explained/?envType=study-plan-v2&envId=leetcode-75
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
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// Case 1: Node do not have any children then we just delete it and nothing else to do here.
		// Case 2: Node has Right child but do not have Left child, just skip this node and link its children with its parent
		if root.Left == nil {
			return root.Right
		}
		// Case 2: Node has Left child but do not have right child, just skip this node and link its children with its parent
		if root.Right == nil {
			return root.Left
		}

		if root.Right != nil && root.Left != nil {
			temp := root.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			root.Val = temp.Val
			root.Right = deleteNode(root.Right, root.Val)
		}
	}

	return root
}

func deleteNodeExercise(root *leetcode.TreeNode, key int) *leetcode.TreeNode {
	return root
}
