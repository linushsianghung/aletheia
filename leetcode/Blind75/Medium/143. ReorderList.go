package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/reorder-list/
/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
*/
func reorderList(head *leetcode.ListNode) {
	slow, fast := head, head
	// Find the middle node of the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the half part after the middle.
	var dummy *leetcode.ListNode
	// Instead of reverse from the slow node, using next node as first node based on the fact that the slow node will be the last node after reordering.
	runner := slow.Next
	for runner != nil {
		temp := runner.Next
		runner.Next = dummy
		dummy = runner
		runner = temp
	}
	// Break the link between 2 new lists
	slow.Next = nil

	runner1, runner2 := head, dummy
	for runner2 != nil {
		temp := runner1.Next
		runner1.Next = runner2
		runner1 = runner1.Next
		runner2 = temp
	}
}
