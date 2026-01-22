package Basic

import "github.com/linushung/aletheia/leetcode"

// ReverseList https://leetcode.com/problems/reverse-linked-list/
/* Given the head of a singly linked list, reverse the list, and return the reversed list. */
func ReverseList(head *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return nil
	}

	/* By using 2 temp pointers (dummy & head) to change the pointer direction. */

	// Cannot use ListNode literal (dummy := &leetcode.ListNode{}) because it will create redundant node like [5,4,3,2,1,0]
	var dummy *leetcode.ListNode

	// Depending on the problem which might need to keep the pointer of the original head
	runner := head
	for runner != nil {
		temp := runner.Next
		/* Reverse the pointer direction */
		runner.Next = dummy
		/* Shift both 2 pointers to each next node*/
		dummy = runner
		runner = temp
	}

	return dummy
}

func reverseListExercise(head *leetcode.ListNode) *leetcode.ListNode {
	return nil
}

func reverseListRecursively(head, previous *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return previous
	}

	temp := head.Next
	head.Next = previous
	return reverseListRecursively(temp, head)
}

// Related Problem: 92. Reverse Linked List II: https://leetcode.com/problems/reverse-linked-list-ii/
// Reference: https://leetcode.com/problems/reverse-linked-list-ii/solutions/30709/talk-is-cheap-show-me-the-code-and-drawing/
func reverseBetween(head *leetcode.ListNode, left int, right int) *leetcode.ListNode {
	dummy := &leetcode.ListNode{Next: head}
	runner := dummy
	for range left - 1 {
		runner = runner.Next
	}

	tail := runner.Next
	for range right - left {
		temp := runner.Next
		runner.Next = tail.Next
		tail.Next = tail.Next.Next
		runner.Next.Next = temp
	}

	return dummy.Next
}
