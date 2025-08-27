package tree

import (
	"github.com/linushung/aletheia/Concepts/Data_Structure/Tree/Tree_Traversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *leetcode.TreeNode) []int { return Tree_Traversal.InorderTraversal(root) }
