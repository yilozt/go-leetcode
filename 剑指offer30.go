package main

type MinStack struct {
	stack, minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] > x {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	this.minStack = this.minStack[:len(this.minStack)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
