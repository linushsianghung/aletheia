package TreeTraversal

import "github.com/linushung/aletheia/leetcode"

// LevelOrder https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/* Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level). */
func LevelOrder(root *leetcode.TreeNode) [][]int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	// return levelOrderRecursively(root)
	return levelOrderIteratively(root)
}

func levelOrderIteratively(root *leetcode.TreeNode) [][]int {
	result := make([][]int, 0)
	queue := []*leetcode.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		nodes := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			// Rather than using "level = append(level, current.Val)" which result in zero leading slices: [[0 3] [0 0 9 20] [0 0 15 7]]
			nodes[i] = current.Val
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		result = append(result, nodes)
	}

	return result
}

// Related Topic: 116. Populating Next Right Pointers in Each Node: https://leetcode.com/problems/populating-next-right-pointers-in-each-node
func connect(root *leetcode.ListNode) *leetcode.ListNode {
	if root == nil {
		return root
	}

	return root
}

// Reference: https://leetcode.com/problems/binary-tree-level-order-traversal/solutions/33445/java-solution-using-dfs/
// Using DFS (Pre-Order Traversal) to realise BFS
func levelOrderRecursively(root *leetcode.TreeNode) [][]int {
	result := make([][]int, 0)

	var localRecursiveFunc func(root *leetcode.TreeNode, level int)
	localRecursiveFunc = func(root *leetcode.TreeNode, level int) {
		if root == nil {
			return
		}

		if len(result) == level {
			result = append(result, make([]int, 0))
		}

		result[level] = append(result[level], root.Val)
		localRecursiveFunc(root.Left, level+1)
		localRecursiveFunc(root.Right, level+1)
	}

	localRecursiveFunc(root, 0)
	return result
}
