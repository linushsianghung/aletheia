package Runners

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/remove-linked-list-elements/
// Ref: https://leetcode.com/problems/remove-linked-list-elements/solutions/1572932/java-three-simple-clean-solutions-w-explanation-iterative-recursive-beats-100/
/*
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new
head.
*/
func removeElements(head *leetcode.ListNode, val int) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return nil
	}

	dummy := &leetcode.ListNode{Next: head}
	runner := dummy

	for runner != nil && runner.Next != nil {
		if runner.Next.Val == val {
			runner.Next = runner.Next.Next
		} else {
			runner = runner.Next
		}
	}

	return dummy.Next
}

func removeElementsExercise(head *leetcode.ListNode, val int) *leetcode.ListNode {

	return nil
}
