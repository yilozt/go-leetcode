package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	stk := make([]int, 0)

	for p := head; p != nil; p = p.Next {
		stk = append(stk, p.Val)
	}

	res := make([]int, len(stk))
	for i := 0; i < len(res); i++ {
		res[i] = stk[len(stk)-1-i]
	}

	return res
}

// end

func main() {
	lst := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	if reflect.DeepEqual([]int{3, 2, 1}, reversePrint(&lst)) {
		fmt.Printf("Done\n")
	}
}
