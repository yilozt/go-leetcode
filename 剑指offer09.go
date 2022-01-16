package main

type stack []int

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) pop() int {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func NewStack() *stack {
	s := stack(make([]int, 0))
	return &s
}

type CQueue struct {
	pos, neg *stack
}

func Constructor() CQueue {
	return CQueue{
		pos: NewStack(),
		neg: NewStack(),
	}
}

func (c *CQueue) AppendTail(item int) {
	c.pos.push(item)
}

func (c *CQueue) DeleteHead() int {
	if c.neg.empty() {
		for !c.pos.empty() {
			c.neg.push(c.pos.pop())
		}
	}
	if c.neg.empty() {
		return -1
	}
	return c.neg.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
