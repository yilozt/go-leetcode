/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
package main

type MyQueue struct {
	forward []int
	reverse []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.reverse) > 0 {
		this.forward = append(this.forward, this.reverse[len(this.reverse)-1])
		this.reverse = this.reverse[:len(this.reverse)-1]
	}
	this.forward = append(this.forward, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for len(this.forward) > 0 {
		this.reverse = append(this.reverse, this.forward[len(this.forward)-1])
		this.forward = this.forward[:len(this.forward)-1]
	}
	result := this.reverse[len(this.reverse)-1]
	this.reverse = this.reverse[:len(this.reverse)-1]
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for len(this.forward) > 0 {
		this.reverse = append(this.reverse, this.forward[len(this.forward)-1])
		this.forward = this.forward[:len(this.forward)-1]
	}
	return this.reverse[len(this.reverse)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.forward) == 0 && len(this.reverse) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
