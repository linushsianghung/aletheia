package CircularQueue

// MyCircularQueue https://leetcode.com/problems/design-circular-queue/
// Reference: https://www.youtube.com/watch?v=8sjFA-IX-Ww
/*
Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle,
and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full,
we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Implement the MyCircularQueue class:ƒ

- MyCircularQueue(k) Initializes the object with the size of the queue to be k.
- int Front() Gets the front item from the queue. If the queue is empty, return -1.
- int Rear() Gets the last item from the queue. If the queue is empty, return -1.
- boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
- boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
- boolean isEmpty() Checks whether the circular queue is empty or not.
- boolean isFull() Checks whether the circular queue is full or not.

You must solve the problem without using the built-in queue data structure in your programming language.
*/
type MyCircularQueue struct {
	buffer     []int
	size       int
	writeIndex int
	readIndex  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		buffer:     make([]int, k),
		size:       k,
		writeIndex: -1,
		readIndex:  -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.readIndex++
	}

	this.writeIndex = (this.writeIndex + 1) % this.size
	this.buffer[this.writeIndex] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.readIndex == this.writeIndex {
		this.writeIndex = -1
		this.readIndex = -1
	} else {
		this.readIndex = (this.readIndex + 1) % this.size
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.buffer[this.readIndex]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.buffer[this.writeIndex]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.writeIndex == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.writeIndex+1)%this.size == this.readIndex
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
