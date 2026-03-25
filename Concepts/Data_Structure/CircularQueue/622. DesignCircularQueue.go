package CircularQueue

// https://leetcode.com/problems/design-circular-queue/
/*
Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle, and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Implement the MyCircularQueue class:
- MyCircularQueue(k) Initializes the object with the size of the queue to be k.
- int Front() Gets the front item from the queue. If the queue is empty, return -1.
- int Rear() Gets the last item from the queue. If the queue is empty, return -1.
- boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
- boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
- boolean isEmpty() Checks whether the circular queue is empty or not.
- boolean isFull() Checks whether the circular queue is full or not.

You must solve the problem without using the built-in queue data structure in your programming language.

Analysis:
Below implementation uses one more variable "count" with front and rear to maintain the internal status of the queue.
The benefit is that the implementation of the functions, like IsEmpty and IsFull, becomes more intuitive and easier to understand.
Also use 0, and -1 for the initial value for front and rear pointers which might not be consistent for both variable, but it exactly describes the current status of the queue:
- front == 0 means if the queue is not empty, just read the element of current index
- rear == -1 means there is no element in the queue
*/

type CircularQueue struct {
	buffer []int // A slice to store the queue's elements.
	front  int   // The index of the front element.
	rear   int   // The index of the last element.
	size   int   // The maximum capacity of the queue.
	count  int   // The current number of elements in the queue.
}

// Construct Initializes the object with the size of the queue to be k.
func Construct(k int) CircularQueue {
	return CircularQueue{
		buffer: make([]int, k),
		front:  0,
		rear:   -1, // Initialize rear to -1 to indicate an empty queue
		size:   k,
		count:  0,
	}
}

// EnQueue Inserts an element into the circular queue. Return true if the operation is successful.
func (q *CircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	q.rear = (q.rear + 1) % q.size
	q.buffer[q.rear] = value
	q.count++

	return true
}

// DeQueue Deletes an element from the circular queue. Return true if the operation is successful.
func (q *CircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.front = (q.front + 1) % q.size
	q.count--

	return true
}

// Front Gets the front item from the queue. If the queue is empty, return -1.
func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.buffer[q.front]
}

// Rear Gets the last item from the queue. If the queue is empty, return -1.
func (q *CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.buffer[q.rear]
}

// IsEmpty Checks whether the circular queue is empty or not.
func (q *CircularQueue) IsEmpty() bool {
	return q.count == 0
}

// IsFull Checks whether the circular queue is full or not.
func (q *CircularQueue) IsFull() bool {
	return q.count == q.size
}
