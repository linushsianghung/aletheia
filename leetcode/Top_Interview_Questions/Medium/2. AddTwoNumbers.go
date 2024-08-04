package Medium

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/add-two-numbers/description/
/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
// Time Complexity: O(n)
func addTwoNumbers(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &leetcode.ListNode{}
	runner := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		d1, d2 := 0, 0

		if l1 != nil {
			d1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			d2 = l2.Val
			l2 = l2.Next
		}

		sum := d1 + d2 + carry
		digit := sum % 10
		carry = sum / 10

		runner.Next = &leetcode.ListNode{Val: digit}
		runner = runner.Next
	}

	return head.Next
}

func addTwoNumbersExercise(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {

	return nil
}
