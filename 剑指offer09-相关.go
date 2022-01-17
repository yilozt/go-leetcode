package main

import "fmt"

type queue []int

func (q *queue) add(item int) {
	*q = append(*q, item)
}

func (q *queue) remove() int {
	if len(*q) == 0 {
		return -1
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *queue) len() int {
	return len(*q)
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func NewQueue() *queue {
	queue := make(queue, 0)
	return &queue
}

type CStack struct {
	pos, neg *queue
}

func Constructor() CStack {
	return CStack{
		pos: NewQueue(),
		neg: NewQueue(),
	}
}

func (c *CStack) AppendTail(item int) {
	if c.pos.empty() {
		c.neg.add(item)
	} else {
		c.pos.add(item)
	}
}

func (c *CStack) DeleteHead() int {
	var oper, another *queue
	if c.pos.empty() {
		oper = c.neg
		another = c.pos
	} else {
		oper = c.pos
		another = c.neg
	}
	for oper.len() > 1 {
		another.add(oper.remove())
	}
	return oper.remove()

}

func main() {
	stack := Constructor()
	stack.AppendTail(1)
	fmt.Println(stack.DeleteHead())
	fmt.Println(stack.DeleteHead())
	stack.AppendTail(2)
	stack.AppendTail(3)
	fmt.Println(stack.DeleteHead())
	stack.AppendTail(4)
	fmt.Println(stack.DeleteHead())
	fmt.Println(stack.DeleteHead())
}
