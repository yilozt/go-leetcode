/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	left, right := head, head
	for k > 1 {
		if right.Next != nil {
			right = right.Next
			k--
		} else {
			return nil
		}
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	return left
}
