package LinkedList

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/intersection-of-two-linked-lists/
/*
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no
intersection at all, return null.
*/
/*
Discussion:
- https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/49785/Java-solution-without-knowing-the-difference-in-len!
- https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/1092898/JS-Python-Java-C%2B%2B-or-Easy-O(1)-Extra-Space-Solution-w-Visual-Explanation

Analysis:
We can use two iterations to do that.
In the first iteration, we will reset the pointer of one linkedlist to the head of another linkedlist after it reaches
the tail node. In the second iteration, we will move two pointers until they points to the same node.
Our operations in the first iteration will help us counteract the difference. So if two linkedlist intersects, the meeting
point in the second iteration must be the intersection point. If the two linked lists have no intersection at all, then the
meeting pointer in the second iteration must be the tail node of both lists, which is null
*/
func getIntersectionNode(headA, headB *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if headA == nil || headB == nil {
		return nil
	}

	runnerA, runnerB := headA, headB
	for runnerA != runnerB {
		// if runnerA.Next == nil => Time Limit Exceeded
		if runnerA == nil {
			runnerA = headB
		} else {
			runnerA = runnerA.Next
		}

		if runnerB == nil {
			runnerB = headA
		} else {
			runnerB = runnerB.Next
		}
	}

	return runnerA // either runnerA == runnerB (a middle node) or runnerA == runnerB == nil (gone through both lists)
}

func getIntersectionNodeExercise(headA, headB *leetcode.ListNode) *leetcode.ListNode {

	return nil
}
