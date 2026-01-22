package Tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/
/*
Given the root of an n-ary tree, return the postorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value
*/
func postorder(root *leetcode.Node) []int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	return postorderRecursively(root)
	// return postorderIteratively(root)
}

func postorderRecursively(root *leetcode.Node) []int {
	result := make([]int, 0)

	var nestedFunc func(node *leetcode.Node)
	nestedFunc = func(node *leetcode.Node) {
		if node == nil {
			return
		}

		for _, child := range node.Children {
			nestedFunc(child)
		}
		result = append(result, node.Val)
	}
	nestedFunc(root)
	return result
}

func postorderIteratively(root *leetcode.Node) []int {
	result := make([]int, 0)
	stack := []*leetcode.Node{root}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{current.Val}, result...)
		stack = append(stack, current.Children...)
	}

	return result
}

func postorderExercise(root *leetcode.Node) []int {
	return nil
}
