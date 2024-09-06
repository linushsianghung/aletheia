package LinkedList

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/odd-even-linked-list/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/odd-even-linked-list/solutions/1606963/c-simplest-solution-w-explanation-one-pass/?envType=study-plan-v2&envId=leetcode-75
/*
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
The first node is considered odd, and the second node is even, and so on.
Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.
*/
func oddEvenList(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	counter := 1
	runner, odd, even := head, head, head.Next

	for runner.Next != nil {
		// Keep the pointer of the next node for runner to go through the linked list
		temp := runner.Next
		// Update the pointer of each node to link to next next node
		runner.Next = runner.Next.Next
		// Move the runner to next node
		runner = temp

		// Move the odd pointer toward the tail so that it can link to even list at the end
		if counter%2 == 1 && odd.Next != nil {
			odd = odd.Next
		}
		counter++
	}

	odd.Next = even
	return head
}
