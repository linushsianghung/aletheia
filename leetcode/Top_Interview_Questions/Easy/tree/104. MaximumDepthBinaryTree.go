package tree

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Top_Interview_Questions"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Ref:
// - https://www.youtube.com/watch?v=hTM3phVI6YQ
// - https://leetcode.com/problems/maximum-depth-of-binary-tree/solutions/1770060/c-recursive-dfs-example-dry-run-well-explained/?envType=featured-list&envId=top-interview-questions?envType=featured-list&envId=top-interview-questions
/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Analysis:
Lets have faith in recursion and assume that we are already given the maximum depth of root's left and right subtrees by recursion.
So to find the maximum depth of this binary tree, we will have to take out the maximum of the 2 depths given to us by recursion, and add 1 to that to consider the current level i.e. root's level into our depth.
```
int maxDepthLeft = maxDepth(root->left);
int maxDepthRight = maxDepth(root->right);
return max(maxDepthLeft, maxDepthRight) + 1;
```
*/
func maxDepth(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}

	return maxDepthDFSRecursively(root)
}

func maxDepthDFSRecursively(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Top_Interview_Questions.Max(maxDepthDFSRecursively(root.Left), maxDepthDFSRecursively(root.Right))
}

func maxDepthDFSRecursivelyExercise(root *leetcode.TreeNode) int {

	return 0
}

// Level Order Traversal
func maxDepthBFSIteratively(root *leetcode.TreeNode) int {
	level := 0
	queue := []*leetcode.TreeNode{root}

	for len(queue) > 0 {
		length := len(queue) // cache the amount of nodes for this level

		for i := 0; i < length; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		level++
	}

	return level
}

func maxDepthBFSIterativelyExercise(root *leetcode.TreeNode) int {

	return 0
}
