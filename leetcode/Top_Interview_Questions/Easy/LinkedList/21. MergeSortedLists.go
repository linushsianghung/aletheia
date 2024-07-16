package LinkedList

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/merge-two-sorted-lists/
/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/
func mergeTwoLists(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	return mergeTwoListsIteratively(l1, l2)
	// return mergeTwoListsRecursively(l1, l2)
}

func mergeTwoListsIteratively(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {

	dummy := &leetcode.ListNode{} /* For aggregating 2 lists into this new node, rather than merging to each other */
	runner := dummy               /* For traversing */

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			runner.Next = l1
			l1 = l1.Next
		} else {
			runner.Next = l2
			l2 = l2.Next
		}

		runner = runner.Next
	}

	if l1 != nil {
		runner.Next = l1
	}

	if l2 != nil {
		runner.Next = l2
	}

	return dummy.Next
}

func mergeTwoListsIterativelyExercise(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	var dummy *leetcode.ListNode
	runner := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			runner.Next = l1
			l1 = l1.Next
		} else {
			runner.Next = l2
			l2 = l2.Next
		}
		runner = runner.Next
	}

	if l1 == nil {
		runner.Next = l2
	} else {
		runner.Next = l1
	}

	return dummy.Next
}

func mergeTwoListsRecursively(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	return nil
}
