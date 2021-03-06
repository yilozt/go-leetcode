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
	if head == nil {
		return nil
	}
	cur := head
	// 拼接链表
	for cur != nil {
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur = cur.Next.Next
	}
	// 调整random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 拆分链表
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
