package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

func createTree(values []int) *TreeNode {
	nodes := make([]*TreeNode, len(values))

	for i, val := range values {
		nodes[i] = &TreeNode{Val: val}
	}

	// Define Tree Structure
	nodes[0].Left = nodes[1]
	nodes[0].Right = nodes[2]
	nodes[2].Left = nodes[3]
	nodes[2].Right = nodes[4]

	return nodes[0]
}
