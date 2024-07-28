package Runners

import (
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
/* Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well. */
func deleteDuplicates(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	runner := head
	for runner != nil && runner.Next != nil {
		if runner.Val == runner.Next.Val {
			runner.Next = runner.Next.Next
		} else {
			runner = runner.Next
		}
	}

	return head
}

func deleteDuplicatesExercise(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	return nil
}
