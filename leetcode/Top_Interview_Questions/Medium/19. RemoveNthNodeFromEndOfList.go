package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description
/* Given the head of a linked list, remove the nth node from the end of the list and return its head. */
func removeNthFromEnd(head *leetcode.ListNode, n int) *leetcode.ListNode {
	dummy := &leetcode.ListNode{Next: head}
	// 2 pointers have to start from the dummy node in case of removing the first node
	slow, fast := dummy, dummy

	// First, move fast so that the gap between slow and fast becomes n
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Move fast and slow to the end to maintain the gap
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Skip the target node
	slow.Next = slow.Next.Next

	return dummy.Next
}

func removeNthFromEndExercise(head *leetcode.ListNode, n int) *leetcode.ListNode {

	return nil
}
