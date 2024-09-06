package data_structure

// MyQueue https://leetcode.com/problems/implement-queue-using-stacks/description/
/*
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/
type MyQueue struct {
	mainStack   []int
	helperStack []int
}

func QueueConstructor() MyQueue {
	return MyQueue{
		mainStack:   make([]int, 0),
		helperStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.mainStack) > 0 {
		this.helperStack = append(this.helperStack, this.mainStack[len(this.mainStack)-1])
		this.mainStack = this.mainStack[:len(this.mainStack)-1]
	}

	this.mainStack = append(this.mainStack, x)

	for len(this.helperStack) > 0 {
		this.mainStack = append(this.mainStack, this.helperStack[len(this.helperStack)-1])
		this.helperStack = this.helperStack[:len(this.helperStack)-1]
	}
}

func (this *MyQueue) Pop() int {
	value := this.mainStack[len(this.mainStack)-1]
	this.mainStack = this.mainStack[:len(this.mainStack)-1]
	return value
}

func (this *MyQueue) Peek() int {
	return this.mainStack[len(this.mainStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.mainStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
