package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]

	len_l := 0
	for i, v := range inorder {
		if v == val {
			len_l = i
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  buildTree(preorder[1:1+len_l], inorder[0:len_l]),
		Right: buildTree(preorder[1+len_l:], inorder[1+len_l:]),
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s[1:1])
}
