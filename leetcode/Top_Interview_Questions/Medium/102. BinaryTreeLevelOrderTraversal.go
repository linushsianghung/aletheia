package Medium

import (
	"github.com/linushung/aletheia/Concepts/Data Structure/Tree/Tree Traversal"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
func levelOrder(root *leetcode.TreeNode) [][]int {
	return Tree_Traversal.LevelOrder(root)
}
