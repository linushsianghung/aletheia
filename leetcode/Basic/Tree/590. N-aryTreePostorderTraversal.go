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

func postorderRecursively(node *leetcode.Node) []int {
	result := make([]int, 0)

	var localRecursiveFunc func(node *leetcode.Node) []int
	localRecursiveFunc = func(node *leetcode.Node) []int {
		for _, node := range node.Children {
			localRecursiveFunc(node)
		}

		result = append(result, node.Val)
		return result
	}

	return localRecursiveFunc(node)
}

func postorderIteratively(node *leetcode.Node) []int {
	result := make([]int, 0)
	stack := []*leetcode.Node{node}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{current.Val}, result...)
		stack = append(stack, current.Children...)
	}

	return result
}

func postorderIterativelyExercise(node *leetcode.Node) []int {

	return nil
}
