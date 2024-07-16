package Medium

import (
	"github.com/linushung/aletheia/Concepts/algorithm/TreeTraversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
func zigzagLevelOrder(root *leetcode.TreeNode) [][]int {
	return TreeTraversal.ZigzagLevelOrder(root)
}
