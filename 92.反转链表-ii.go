/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var prev, next *ListNode
	var dummy = &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	for i := 1; i < m; i++ {
		head = head.Next
	}
	begin := head
	end := head.Next
	head = head.Next
	for i := m; i <= n; i++ {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	begin.Next = prev
	end.Next = head
	return dummy.Next
}

// @lc code=end
