package tree

import (
	TwoPointers "github.com/linushung/aletheia/Concepts/algorithm/BinarySearch"
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Ref:
// - https://www.youtube.com/watch?v=0K0uCMYq5ng
// - https://www.youtube.com/watch?v=12omz-VAyRk
func sortedArrayToBST(nums []int) *leetcode.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// i := len(nums) / 2
	// root := &leetcode.TreeNode{Val: nums[i]}
	// root.Left = sortedArrayToBST(nums[:i])
	// root.Right = sortedArrayToBST(nums[i + 1:])
	// return root

	return constructSubBST(nums, 0, len(nums)-1)
}

func constructSubBST(nums []int, left, right int) *leetcode.TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &leetcode.TreeNode{Val: nums[mid]}
	node.Left = constructSubBST(nums, left, mid-1)
	node.Right = constructSubBST(nums, mid+1, right)

	return node
}

// Related Topic: 704 - Binary Search https://leetcode.com/problems/binary-search/description/
func search(nums []int, target int) int {
	return TwoPointers.Search(nums, target)
}
