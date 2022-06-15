package main

import "fmt"

type Stk []int

func (this *Stk) Push(value int) { *this = append(*this, value) }
func (this *Stk) Pop() int {
	t := len(*this) - 1
	tail := (*this)[t]
	*this = (*this)[0:t]
	return tail
}

type CQueue struct {
	in  Stk
	out Stk
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	out_len := len(this.out)

	if out_len != 0 {
	} else {
		in_len := len(this.in)
		for i := 0; i < in_len; i++ {
			this.out.Push(this.in.Pop())
		}

	}

	if len(this.out) != 0 {
		return this.out.Pop()
	} else {
		return -1
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	q := CQueue{}

	q.AppendTail(1)
	q.AppendTail(2)
	q.AppendTail(3)

	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}
