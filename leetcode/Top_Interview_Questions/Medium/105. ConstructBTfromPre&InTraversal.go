package Medium

import (
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree,
construct and return the binary tree.

Analysis 1:
For 2 arrays, preorder and inorder, preorder traversing implies that preorder[0] is the root node. Then we can find this preorder[0] in the inorder,
say it's inorder[5]. Now we know that inorder[5] is root, so that inorder[0] ~ inorder[4] is on the left side and inorder[6] to the end is on the right side.
Recursively doing this on sub-arrays to build a tree out of it.

Analysis 2 (How to get the index of the right child):
- Our aim is to find out the index of right child for current node in the preorder array
- We know the index of current node in the preorder array - preStart (or whatever you call it), it's the root of a subtree
- Remember preorder traversal always visit all the nodes on left branch before going to the right ( root -> left -> ... -> right),
  therefore, we can get the immediate right child index by skipping all the node on the left branches/subtrees of current node
- The inorder array has this information exactly. Remember when we found the root in "inorder" array, we immediately know how many nodes
  are on the left subtree and how many are on the right subtree
- Therefore the immediate right child index is preStart + numsOnLeft + 1 (remember in preorder traversal array root is always ahead of
  children nodes, but you don't know which one is the left child which one is the right, and this is why we need inorder array)
- numsOnLeft = root - inStart.
*/
func buildTree(preorder []int, inorder []int) *leetcode.TreeNode {

	return buildTreeRecursively(0, 0, len(inorder)-1, preorder, inorder)
}

func buildTreeRecursively(preStart int, inStart int, inEnd int, preorder []int, inorder []int) *leetcode.TreeNode {
	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}

	root := &leetcode.TreeNode{Val: preorder[preStart]}
	idx := 0
	for i, in := range inorder {
		if in == root.Val {
			idx = i
			break
		}
	}

	root.Left = buildTreeRecursively(preStart+1, inStart, idx-1, preorder, inorder)
	root.Right = buildTreeRecursively(preStart+idx-inStart+1, idx+1, inEnd, preorder, inorder)
	return root
}
