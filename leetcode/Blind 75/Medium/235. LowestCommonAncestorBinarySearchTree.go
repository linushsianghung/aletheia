package Medium

import (
	"github.com/linushung/aletheia/Concepts/Data Structure/Tree"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree
func lowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	return Tree.LowestCommonAncestor(root, p, q)
}
