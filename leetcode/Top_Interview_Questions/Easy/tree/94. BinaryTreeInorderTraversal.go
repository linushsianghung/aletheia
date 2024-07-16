package tree

import (
	"github.com/linushung/aletheia/Concepts/algorithm/TreeTraversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *leetcode.TreeNode) []int { return TreeTraversal.InorderTraversal(root) }
