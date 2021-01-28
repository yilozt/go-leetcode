/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
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

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeList(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for left != nil && right != nil {
		if left.Val > right.Val {
			current.Next = right
			right = right.Next
		} else {
			current.Next = left
			left = left.Next
		}
		current = current.Next
	}
	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}
	return dummy.Next
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	return mergeList(left, right)
}

func findMiddle(head *ListNode) *ListNode {
	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}

// @lc code=end
