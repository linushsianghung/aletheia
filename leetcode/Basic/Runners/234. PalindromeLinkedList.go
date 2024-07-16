package Runners

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Basic"
)

// https://leetcode.com/problems/palindrome-linked-list/
/* Given the head of a singly linked list, return true if it is palindrome or false otherwise. */
/*
Discussion:
- https://leetcode.com/problems/palindrome-linked-list/solutions/1137027/js-python-java-c-easy-floyd-s-reversal-solution-w-explanation/
*/
func isPalindromeLinkedList(head *leetcode.ListNode) bool {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return true
	}

	return isPalindromeIteratively(head)
}

func isPalindromeIteratively(head *leetcode.ListNode) bool {
	/* Floyd Cycle Detection Algorithm */
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	runner := Basic.ReverseList(slow)
	for runner != nil {
		if runner.Val != head.Val {
			return false
		}
		head = head.Next
		runner = runner.Next
	}

	return true
}

func isPalindromeIterativelyExercise(head *leetcode.ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	runner := Basic.ReverseList(slow)

	for runner != head && runner.Next != head.Next {
		if runner.Val != head.Val {
			return false
		}
		runner = runner.Next
		head = head.Next
	}

	return false
}

func isPalindromeRecursively(head *leetcode.ListNode) bool {
	return false
}

// Related Topic: 876 - Middle of the Linked List: https://leetcode.com/problems/middle-of-the-linked-list/
func findMiddleNode(head *leetcode.ListNode) *leetcode.ListNode {
	return middleNode(head)
}
