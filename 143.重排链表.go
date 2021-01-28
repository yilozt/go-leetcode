/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil
	tail = reverseList(tail)
	head = mergeList(head, tail)
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	flag := true
	for left != nil && right != nil {
		if flag {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		flag = !flag
		current = current.Next
	}
	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}
	return dummy.Next
}

func findMiddle(head *ListNode) *ListNode {
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// @lc code=end
