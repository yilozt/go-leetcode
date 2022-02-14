package main

type ListNode struct {
	Next *ListNode
}

func LengthOfLoop(head *ListNode) int {
	if head == nil {
		return -1
	}
	slow, fast := head, head
	var steps int
	for slow != nil && fast != nil { // 慢走一，快走二，每走一次领先一步，如果两节点相同，环节点个数等于走的次数
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		steps++
		if slow == fast {
			return steps
		}
	}
	return -1
}
func EntryNodeOfLoop(head *ListNode) *ListNode {
	var length int
	if length = LengthOfLoop(head); length == -1 { // 不存在环
		return nil
	}
	left, right := head, head
	for length >= 1 { // 首先先走环节点个数
		right = right.Next
		length--
	}
	for left != right { // 继续走直到两点相同，此时相遇点是环入口
		left = left.Next
		right = right.Next
	}
	return left
}
