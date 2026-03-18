package Medium

import (
	"github.com/linushung/aletheia/leetcode"
)

// https://leetcode.com/problems/reorder-list/
/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Analysis:
The algorithm is essentially a "Three-Step Recipe":
1. Find the Middle: Split the list into two halves.
2. Reverse the Second Half: Turn the right half backwards so we can iterate it from the end.
3. Merge (Zip): Weave the two lists together
*/
func reorderList(head *leetcode.ListNode) {
	slow, fast := head, head
	// Find the middle node of the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the half part after the middle.
	var prev *leetcode.ListNode
	// Instead of reverse from the slow node, using next node as first node based on the fact that the slow node will be the last node after reordering.
	curr := slow.Next
	// Break the link between the two halves immediately to avoid confusion later
	slow.Next = nil

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	/*
		Merge the two lists: 'head' (first half) and 'prev' (reversed second half)
		Let's look at the new merge loop with an example:
			List 1 (first): 1 -> 2
			List 2 (second): 4 -> 3

		Iteration:
		1. Save Next: tmp1 = 2, tmp2 = 3.
		2. Link 1 to 4: first.Next = second (Result: 1 -> 4).
		3. Link 4 to 2: second.Next = tmp1 (Result: 1 -> 4 -> 2).
		4. Advance: first becomes 2, second becomes 3.

		This explicit step-by-step saving of Next pointers is much safer and easier to debug than the implicit swapping method.
	*/
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func reorderListExercise(head *leetcode.ListNode) {
}

func reorderListAlternative(head *leetcode.ListNode) {
	slow, fast := head, head
	// Find the middle of the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the half part after the middle
	var dummy *leetcode.ListNode
	runner := slow.Next
	for runner != nil {
		temp := runner.Next
		runner.Next = dummy
		dummy = runner
		runner = temp
	}
	slow.Next = nil

	runner = head
	for dummy != nil {
		temp := runner.Next
		runner.Next = dummy
		runner = runner.Next
		dummy = temp
	}
}
