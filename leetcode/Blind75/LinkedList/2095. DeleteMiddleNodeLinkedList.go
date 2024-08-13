package LinkedList

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list
/*
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
*/
func deleteMiddle(head *leetcode.ListNode) *leetcode.ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast, previous := head, head, &leetcode.ListNode{Next: head}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		previous = previous.Next
	}

	previous.Next = previous.Next.Next
	return head
}
