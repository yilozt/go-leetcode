package main

type CQueue struct {
	pos, neg []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// 把反向栈里的内容重新压回来
	for len(this.neg) > 0 {
		this.pos = append(this.pos, this.neg[len(this.neg)-1])
		this.neg = this.neg[:len(this.neg)-1]
	}
	this.pos = append(this.pos, value)
}

func (this *CQueue) DeleteHead() int {
	for len(this.pos) > 0 {
		this.neg = append(this.neg, this.pos[len(this.pos)-1])
		this.pos = this.pos[:len(this.pos)-1]
	}
	if len(this.neg) > 0 {
		result := this.neg[len(this.neg)-1]
		this.neg = this.neg[:len(this.neg)-1]
		return result
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
