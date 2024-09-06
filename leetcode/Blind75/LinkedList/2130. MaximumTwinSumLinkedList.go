package LinkedList

import (
	"github.com/linushung/aletheia/leetcode"
	"github.com/linushung/aletheia/leetcode/Basic"
)

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/description/?envId=leetcode-75
/*
In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.

Given the head of a linked list with even length, return the maximum twin sum of the linked list.
*/
func pairSum(head *leetcode.ListNode) int {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	runner := Basic.ReverseList(slow)

	maxSum := 0
	for runner != nil {
		sum := head.Val + runner.Val
		if maxSum < sum {
			maxSum = sum
		}

		head = head.Next
		runner = runner.Next
	}

	return maxSum
}
