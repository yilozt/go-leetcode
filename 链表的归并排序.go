package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	test := &ListNode{
		val: 1,
		next: &ListNode{
			val: 3,
			next: &ListNode{
				val: 2,
				next: &ListNode{
					val:  4,
					next: nil,
				},
			},
		},
	}
	root := sort(test, nil)
	for root != nil {
		fmt.Println(root.val)
		root = root.next
	}
}

func sort(start, end *ListNode) *ListNode {
	if start == nil {
		return start
	}
	if start.next == end { // ？
		start.next = nil
		return start
	}
	slow, fast := start, start
	for fast != end {
		slow = slow.next
		fast = fast.next
		if fast != end {
			fast = fast.next
		}
	}
	mid := slow
	left := sort(start, mid)
	right := sort(mid, end)
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for left != nil && right != nil {
		if left.val > right.val {
			current.next = right
			right = right.next
		} else {
			current.next = left
			left = left.next
		}
		current = current.next
	}
	if left == nil {
		current.next = right
	} else {
		current.next = left
	}
	return dummy.next
}
