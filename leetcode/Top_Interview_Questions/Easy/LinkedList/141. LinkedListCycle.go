package LinkedList

import "github.com/linushung/aletheia/leetcode"

// https://leetcode.com/problems/linked-list-cycle/
/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.

Analysis: Floyd's cycle-finding algorithm
Ref: https://matthung0807.blogspot.com/2020/04/floyd-cycle-detection-algorithm-floyd.html

For this problem, let's see what will happen if there's a circle.
If it's a little abstract, then let's think about we’re running on a circular track.

If the track is 100m long, your speed is 10m/s, your friend's speed is 5m/s. Then after 20s, you've run 200m, your
friend has run 100m. But because the track is circular, so you will be at the same place with your friend since the
difference between your distances is exactly 100m.
How about we change another track, change the speed of you and your friend? As long as your speeds are different,
the faster person will always catch up with the slower person again.
*/
func hasCycle(head *leetcode.ListNode) bool {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return false
	}

	return hasCycle1(head)
}

func hasCycle1(head *leetcode.ListNode) bool {
	slow, fast := head, head

	// Alternative: fast.Next != nil && fast.Next.Next != nil,
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func hasCycle1Exercise(head *leetcode.ListNode) bool {

	return false
}

/* Ref: https://www.youtube.com/watch?v=KoVAJlSOpXw */
func hasCycle2(head *leetcode.ListNode) bool {

	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
