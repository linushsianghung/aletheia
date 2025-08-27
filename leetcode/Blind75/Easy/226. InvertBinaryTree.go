package Easy

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Basic/Tree"
)

// https://leetcode.com/problems/invert-binary-tree/
func invertTree(root *leetcode.TreeNode) *leetcode.TreeNode {
	return Tree.InvertTree(root)
}
