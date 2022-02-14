package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var mergeHead *ListNode
	if l1.Val > l2.Val {
		mergeHead = l2
		mergeHead.Next = mergeTwoLists(l1, l2.Next)
	} else {
		mergeHead = l1
		mergeHead.Next = mergeTwoLists(l1.Next, l2)
	}
	return mergeHead
}
