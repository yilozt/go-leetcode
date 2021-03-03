/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	left, right := head, head
	for k > 1 {
		right = right.Next
		k--
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	return left
}
