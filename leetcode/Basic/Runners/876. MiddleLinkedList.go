package Runners

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/middle-of-the-linked-list/
/* Ref:
When traversing the list with a pointer slow, make another pointer fast that traverses twice as fast. When fast
reaches the end of the list, slow must be in the middle.

Complexity Analysis:
- Time Complexity: O(N), where N is the number of nodes in the given list.
- Space Complexity: O(1), the space used by slow and fast.
*/
/*
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

Analysis:
1 middle nodes: __ -> __ -> __ -> _s_ -> __ -> __ -> _f_
2 middle nodes: __ -> __ -> __ -> __ -> _s_ -> __ -> __ -> __ -> f
*/
func middleNode(head *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return nil
	}

	slow, fast := head, head
	// For fast.Next != nil && fast.Next.Next != nil
	// If there are two middle nodes, will return the first one
	for fast != nil && fast.Next != nil {
		slow = slow.Next /* Basic traversing operation (go to next node) */
		fast = fast.Next.Next
	}

	return slow
}

func middleNodeExercise(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	return nil
}
