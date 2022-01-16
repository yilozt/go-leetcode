/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reversePrint(head *ListNode) []int {
	var result []int
	PrintRecursively(head, &result)
	return result
}

func PrintRecursively(head *ListNode, result *[]int) {
	if head != nil {
		PrintRecursively(head.Next, result)
		*result = append(*result, head.Val)
	}
}
