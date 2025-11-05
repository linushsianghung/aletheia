package CircularQueue

// https://leetcode.com/problems/design-circular-queue/
/*
Analysis design and invariants:
- This implementation uses exactly two indices (rear and front) and relies on a sentinel value of -1 to represent the empty state.
- Empty is defined as: rear == -1. Note this is asymmetric (it ignores front). It works here only because when the queue becomes empty, both rear and front are reset to -1 together.
- Full is defined as: (rear+1)%size == front. This requires front to be a valid index (0..size-1) whenever the queue isn’t empty.
- Because of the sentinel empty state, EnQueue and DeQueue both need special-case handling (see *(1) and *(2) below). This is the source of the unintuitiveness.

Why it feels unintuitive and where risks might be:
1) IsEmpty uses only rear (-1) to determine emptiness. If future changes ever reset only front to -1 (but not rear), IsEmpty would lie and Front()/Rear() could index out-of-range. Keeping both pointers consistent is critical.
2) EnQueue updates front when inserting into an empty queue (*(1)). Conceptually enQueue shouldn’t touch the read pointer, but here it must, because the empty-state sentinel (-1,-1) requires front to become a valid index before writes proceed.
3) DeQueue special-cases the last element (*(2)) to reset both indices to -1. If this reset is forgotten or done partially, the empty-state check breaks.
4) Edge case: k == 0 would cause modulo-by-zero in IsFull, EnQueue, and DeQueue. LeetCode typically guarantees k >= 1, but this implementation assumes that precondition.

Takeaways:
- The algorithm is correct under its invariants, but the invariants themselves (using -1 sentinels and an asymmetric IsEmpty) force extra conditionals and make the code harder to reason about.
- For a more intuitive design, prefer either (a) tracking a separate count, or (b) initializing front and rear to 0 and representing empty/full via count or a boolean flag. See DesignCircularQueue.go for a count-based version that avoids these special cases.
*/

type MyCircularQueue struct {
	buffer []int
	rear   int // or writeIndex
	front  int // or readIndex
	size   int // For the maximum capacity of the queue, which can be replaced by cap(buffer)
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		buffer: make([]int, k),
		size:   k,
		rear:   -1,
		front:  -1,
	}
}

func (cq *MyCircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		return false
	}

	// *(1): If the queue is empty, need to update the front. It's not quite intuitive because EnQueue should have nothing to do with front
	if cq.IsEmpty() {
		cq.front++
	}

	cq.rear = (cq.rear + 1) % cq.size
	cq.buffer[cq.rear] = value
	return true
}

func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		return false
	}

	// *(2): It means when there is only one element in the queue, reset the both index to initial value. It seems it's necessary to use rear to define empty status based on cq implementation, instead of using 2 pointers. Again it's not quite intuitive.
	if cq.front == cq.rear {
		cq.rear = -1
		cq.front = -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}

	return true
}

func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}

	return cq.buffer[cq.front]
}

func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}

	return cq.buffer[cq.rear]
}

func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.rear == -1
}

func (cq *MyCircularQueue) IsFull() bool {
	return (cq.rear+1)%cq.size == cq.front
}
