package main

import "fmt"

// O（1）删除链表节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head, toBeDeleted *ListNode) *ListNode {
	if head == nil || toBeDeleted == nil {
		return nil
	}
	if head == toBeDeleted { // 要删除的点是头节点
		return head.Next
	}
	if toBeDeleted.Next != nil { // 要删除的节点有下个节点
		toBeDeleted.Val = toBeDeleted.Next.Val   // 将该节点值赋为下个节点
		toBeDeleted.Next = toBeDeleted.Next.Next // 删除下个节点
		return head
	}
	current := head
	for current.Next != nil && current.Next != toBeDeleted {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}

func printListNode(root *ListNode) {
	var list []int
	for root != nil {
		list = append(list, root.Val)
		root = root.Next
	}
	fmt.Println(list)
}

func main() {
	test := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	toBeDeleted := []*ListNode{test, test.Next.Next, test.Next.Next.Next.Next}
	for _, item := range toBeDeleted {
		test = deleteNode(test, item)
		printListNode(test)
	}
}
