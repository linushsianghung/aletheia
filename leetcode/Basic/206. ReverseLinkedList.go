package Basic

import (
	"github.com/linushung/aletheia/leetcode"
)

// ReverseList https://leetcode.com/problems/reverse-linked-list/
/* Given the head of a singly linked list, reverse the list, and return the reversed list. */
func ReverseList(head *leetcode.ListNode) *leetcode.ListNode {
	/* Basic init check for LinkedList and Tree */
	if head == nil {
		return nil
	}

	/* By using 2 temp pointers (prev & head) to change the pointer direction. */

	// Cannot use ListNode literal (prev := &leetcode.ListNode{}) because it will create redundant node like [5,4,3,2,1,0]
	var prev *leetcode.ListNode

	// Depending on the problem which might need to keep the pointer of the original head
	runner := head
	for runner != nil {
		temp := runner.Next
		/* Reverse the pointer direction */
		runner.Next = prev
		/* Shift both 2 pointers to each next node*/
		prev = runner
		runner = temp
	}

	return prev
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

// Related Problem: 92. Reverse Linked List II: https://leetcode.com/problems/reverse-linked-list-ii/
// Reference: https://leetcode.com/problems/reverse-linked-list-ii/solutions/30709/talk-is-cheap-show-me-the-code-and-drawing/
func reverseBetween(head *leetcode.ListNode, left int, right int) *leetcode.ListNode {
	dummy := &leetcode.ListNode{Next: head}
	runner := dummy
	for range left - 1 {
		runner = runner.Next
	}

	// It's necessary to keep the tail pointer
	tail := runner.Next
	for range right - left {
		temp := runner.Next
		// Because the tail pointer, so the runner can always get the next reversed node
		runner.Next = tail.Next
		tail.Next = runner.Next.Next
		runner.Next.Next = temp
	}

	return dummy.Next
}

func reverseBetweenExercise(head *leetcode.ListNode, left int, right int) *leetcode.ListNode {
	return nil
}

/*
# Analysis:
1. Concept Introduction / Explanation
**Analogy:** Imagine a line of people where each person is holding the shoulder of the person in front of them (a singly linked list). To reverse the line so the last person becomes the first, each person must let go of the shoulder in front and instead grab the shoulder of the person who was behind them.

To do this without losing the line, you need a "temporary" person (a pointer) to remember who was originally in front of you before you turn around.

**Core Logic (Iterative):**
We maintain three pointers:
- `prev`: The node that will become the new "next".
- `curr`: The node we are currently processing.
- `nextTemp`: A temporary placeholder to store the next node before we break the link.

**Core Logic (Reverse Between):**
The "insertion" method is often the most efficient for sub-segments. Instead of reversing the whole sub-segment and re-attaching, we take the node *after* the current tail and move it to the front of the sub-segment, one by one.

# 2. Complexity Analysis
**ReverseList (Iterative):**
- **Time Complexity:** O(n), where n is the number of nodes in the list. We visit each node exactly once.
- **Space Complexity:** O(1), as we only use a few pointer variables regardless of the list size.

**ReverseList (Recursive):**
- **Time Complexity:** O(n).
- **Space Complexity:** O(n), due to the implicit recursion stack. If the list is very long, this could lead to a StackOverflowError.

**ReverseList II (Reverse Between):**
- **Time Complexity:** O(n). In the worst case, we traverse the entire list once.
- **Space Complexity:** O(1). We perform the reversal in-place using a dummy node and a few pointers.

# 3. Interview Suggestions
- **Draw it out:** Linked list problems are notoriously hard to track mentally. Draw the nodes and arrows on a whiteboard.
- **Edge Cases:** Always ask/check: Is the head null? Is there only one node? For `reverseBetween`, what if `left == right`?
- **The Dummy Node Trick:** For problems where the head might change (like `reverseBetween` where `left` could be 1), always use a `dummy` node pointing to `head`. It simplifies the logic significantly because you don't need a special case for the head.
- **Identify the "Break":** In an interview, explain *why* you need the `temp` variable (because setting `curr.next = prev` overwrites the only link to the rest of the list).

### Java Code:

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class Solution {

    public ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode curr = head;
        while (curr != null) {
            ListNode nextTemp = curr.next; // Store next
            curr.next = prev;              // Reverse pointer
            prev = curr;                   // Move prev forward
            curr = nextTemp;               // Move curr forward
        }
        return prev;
    }

    public ListNode reverseListRecursive(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode p = reverseListRecursive(head.next);
        head.next.next = head;
        head.next = null;
        return p;
    }

    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (head == null) return null;

        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode pre = dummy;

        // 1. Move 'pre' to the node right before the sub-list starts
        for (int i = 0; i < left - 1; i++) {
            pre = pre.next;
        }

        // 2. 'start' is the first node of the sub-list
        // 'then' is the node that will be moved
        ListNode start = pre.next;
        ListNode then = start.next;

        // 3. Perform the "lifting" of 'then' to the position after 'pre'
        // Example: 1 -> 2 -> 3 -> 4, left=2, right=4
        // pre=1, start=2, then=3
        for (int i = 0; i < right - left; i++) {
            start.next = then.next;
            then.next = pre.next;
            pre.next = then;
            then = start.next;
        }

        return dummy.next;
    }
}
*/
