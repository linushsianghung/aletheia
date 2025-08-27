package Easy

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Basic/Tree"
)

// https://leetcode.com/problems/same-tree/
func isSameTree(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {
	return Tree.IsSameTree(p, q)
}
