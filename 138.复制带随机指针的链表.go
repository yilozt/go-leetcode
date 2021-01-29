/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
package main

func copyRandomList(head *Node) *Node {
	var next *Node
	current := head
	// 交叉链表
	for current != nil {
		next = current.Next
		current.Next = &Node{Val: current.Val, Next: next, Random: nil}
		current = next
	}
	// 复制随机指针
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}
	// 拆分链表（注意不能只返回结果，还需要将传入的链表恢复为初始状态）
	current = head
	dummy := &Node{}
	tmp := dummy
	for current != nil {
		tmp.Next = current.Next
		tmp = tmp.Next
		current.Next = current.Next.Next
		current = current.Next
	}
	return dummy.Next
}

// @lc code=end
