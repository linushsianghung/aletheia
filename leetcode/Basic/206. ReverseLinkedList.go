package Basic

import "github.com/linushung/aletheia/leetcode"

// ReverseList https://leetcode.com/problems/reverse-linked-list/
/* Given the head of a singly linked list, reverse the list, and return the reversed list. */
func ReverseList(head *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return nil
	}

	/*
		By using 2 temp pointers (dummy & next) to change the pointer direction.
		The basic idea is "Pointer1 = Pointer2" which means that: Make Pointer1 points to the Node of Pointer2
	*/
	var dummy *leetcode.ListNode
	for head != nil {
		temp := head.Next
		/* Reverse the pointer direction */
		head.Next = dummy
		/* Shift both 2 pointers to each next node*/
		dummy = head
		head = temp
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
