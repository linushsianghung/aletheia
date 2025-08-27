package Medium

import (
	"github.com/linushung/aletheia/Concepts/Data_Structure/Tree/Tree_Traversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *leetcode.TreeNode) [][]int {
	return Tree_Traversal.LevelOrder(root)
}
