package Medium

import (
	"github.com/linushung/aletheia/Concepts/Data_Structure/Tree/Tree_Traversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
func zigzagLevelOrder(root *leetcode.TreeNode) [][]int {
	return Tree_Traversal.ZigzagLevelOrder(root)
}
