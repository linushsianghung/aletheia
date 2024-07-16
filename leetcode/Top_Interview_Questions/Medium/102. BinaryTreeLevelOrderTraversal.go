package Medium

import (
	"github.com/linushung/aletheia/Concepts/algorithm/TreeTraversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
func levelOrder(root *leetcode.TreeNode) [][]int {
	return TreeTraversal.LevelOrder(root)
}
