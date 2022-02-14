package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	var vals []int
	for current := head; current != nil; current = current.Next {
		vals = append(vals, current.Val)
	}
	fmt.Println(vals)
}

func main() {
	test := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	printList(test)
	printList(reverseList(test))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	prev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return prev
}
