package data_structure

// MyStack https://leetcode.com/problems/implement-stack-using-queues/description/
/*
Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:

void push(int x) Pushes element x to the top of the stack.
int pop() Removes the element on the top of the stack and returns it.
int top() Returns the element on the top of the stack.
boolean empty() Returns true if the stack is empty, false otherwise.
Notes:

You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.
*/
type MyStack struct {
	mainQueue   []int
	helperQueue []int
}

func StackConstructor() MyStack {
	return MyStack{
		mainQueue:   make([]int, 0),
		helperQueue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	for len(this.mainQueue) > 0 {
		this.helperQueue = append(this.helperQueue, this.mainQueue[0])
		this.mainQueue = this.mainQueue[1:]
	}

	this.mainQueue = append(this.mainQueue, x)

	for len(this.helperQueue) > 0 {
		this.mainQueue = append(this.mainQueue, this.helperQueue[0])
		this.helperQueue = this.helperQueue[1:]
	}
}

func (this *MyStack) Pop() int {
	value := this.mainQueue[0]
	this.mainQueue = this.mainQueue[1:]
	return value
}

func (this *MyStack) Top() int {
	return this.mainQueue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.mainQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
