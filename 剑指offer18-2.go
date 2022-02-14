package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	test := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	printListNode(deleteDuplication(test))
}

func printListNode(root *ListNode) {
	var list []int
	for root != nil {
		list = append(list, root.Val)
		root = root.Next
	}
	fmt.Println(list)
}

func deleteDuplication(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			val := current.Next.Val
			for current.Next != nil && current.Next.Val == val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}
