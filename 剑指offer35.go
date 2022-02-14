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
	// 一种空间复杂度O1，时间复杂度On的算法
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil { // 将复制后的节点放在节点之后，进行拼接
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur = cur.Next.Next
	}
	// 调整random指针
	cur = head
	for cur != nil { // random指针的下个节点就是复制后的random节点
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 最后再将链表拆分，保留原始链表的前提下拆出复制链表
	result := head.Next
	cur = head.Next
	prev := head
	for cur.Next != nil {
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev = prev.Next
		cur = cur.Next
	}
	prev.Next = nil
	return result
}
