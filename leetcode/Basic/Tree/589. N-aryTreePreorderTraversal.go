package Tree

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
/*
Given the root of an n-ary tree, return the preorder traversal of its currents' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value
*/
func preorder(root *leetcode.Node) []int {
	/* Basic init check for LinkedList and Tree */
	if root == nil {
		return nil
	}

	return preorderRecursively(root)
	// return preorderIteratively(root)
}

func preorderRecursively(root *leetcode.Node) []int {
	result := make([]int, 0)

	var localRecursiveFunc func(root *leetcode.Node)
	localRecursiveFunc = func(root *leetcode.Node) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		for _, child := range root.Children {
			localRecursiveFunc(child)
		}
	}

	localRecursiveFunc(root)
	return result
}

func preorderIteratively(root *leetcode.Node) []int {
	result := make([]int, 0)
	stack := []*leetcode.Node{root}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Based on the LeetCode example, the Left subtree has to be processed first, so reverse the children nodes
		reverse(current.Children)
		for _, child := range current.Children {
			stack = append(stack, child)
		}
	}

	return result
}

func preorderIterativelyExercise(root *leetcode.Node) []int {

	return nil
}

func reverse(slice []*leetcode.Node) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
