/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
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

func partition(head *ListNode, x int) *ListNode {
	largeDummy, smallDummy := &ListNode{}, &ListNode{}
	largeCurrent, smallCurrent := largeDummy, smallDummy
	for head != nil {
		if head.Val < x {
			smallCurrent.Next = head
			smallCurrent = smallCurrent.Next
		} else {
			largeCurrent.Next = head
			largeCurrent = largeCurrent.Next
		}
		head = head.Next
	}
	smallCurrent.Next = largeDummy.Next
	// 不加这行可能导致出现环进而OOM
	largeCurrent.Next = nil
	return smallDummy.Next
}

// @lc code=end
