package Medium

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Top_Interview_Questions/Medium"
)

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func buildTree(preorder []int, inorder []int) *leetcode.TreeNode {
	return Medium.BuildTree(preorder, inorder)
}
