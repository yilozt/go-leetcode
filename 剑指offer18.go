/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			// 因为链表中节点的值互不相同，所以可以直接break
			break
		} else {
			head = head.Next
		}
	}
	return dummy.Next

}
