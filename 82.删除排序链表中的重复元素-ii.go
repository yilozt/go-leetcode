/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	lst := make([]*ListNode, 0)

	p := head
	for p != nil {
		// 查询重复节点
		repeat := false
		for p.Next != nil && p.Next.Val == p.Val {
			repeat = true
			p = p.Next
		}

		if !repeat {
			lst = append(lst, p)
		}

		tmp := p.Next
		p.Next = nil
		p = tmp
	}

	len := len(lst)
	if len == 0 {
		return nil
	} else {
		for i := 0; i < len-1; i++ {
			lst[i].Next = lst[i+1]
		}
		return lst[0]
	}
}

// @lc code=end
