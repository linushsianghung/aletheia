package Basic

import (
	"github.com/linushung/aletheia/leetcode"
)

// ReverseList https://leetcode.com/problems/reverse-linked-list/
/* Given the head of a singly linked list, reverse the list, and return the reversed list. */
func ReverseList(head *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return nil
	}

	/* By using 2 temp pointers (prev & head) to change the pointer direction. */

	// Cannot use ListNode literal (prev := &leetcode.ListNode{}) because it will create redundant node like [5,4,3,2,1,0]
	var prev *leetcode.ListNode

	// Depending on the problem which might need to keep the pointer of the original head
	runner := head
	for runner != nil {
		temp := runner.Next
		/* Reverse the pointer direction */
		runner.Next = prev
		/* Shift both 2 pointers to each next node*/
		prev = runner
		runner = temp
	}

	return prev
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

	// It's necessary to keep the tail pointer
	tail := runner.Next
	for range right - left {
		temp := runner.Next
		// Because the tail pointer, so the runner can always get the next reversed node
		runner.Next = tail.Next
		tail.Next = runner.Next.Next
		runner.Next.Next = temp
	}

	return dummy.Next
}

func reverseBetweenExercise(head *leetcode.ListNode, left int, right int) *leetcode.ListNode {
	dummy := &leetcode.ListNode{Next: head}
	runner := dummy

	for range left - 1 {
		runner = runner.Next
	}

	tail := runner.Next
	for range right - left {
		temp := runner.Next
		runner.Next = tail.Next
		tail.Next = runner.Next.Next
		runner.Next.Next = temp
	}

	return nil
}
