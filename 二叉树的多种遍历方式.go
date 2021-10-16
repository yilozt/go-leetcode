package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func main() {
	tree := &TreeNode{
		val:   1,
		left:  &TreeNode{val: 2, left: &TreeNode{val: 4}, right: &TreeNode{val: 5}},
		right: &TreeNode{val: 3, left: &TreeNode{val: 6}, right: &TreeNode{val: 7}},
	}
	preorder(tree)
	inorder(tree)
	postorder(tree)
}

func preorder(root *TreeNode) {
	var (
		stack  []*TreeNode
		result []int
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.val)
			stack = append(stack, root)
			root = root.left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.right
	}
	fmt.Println(result)
}

func inorder(root *TreeNode) {
	var (
		stack  []*TreeNode
		result []int
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.val)
		root = node.right
	}
	fmt.Println(result)
}

func postorder(root *TreeNode) {
	var (
		stack     []*TreeNode
		result    []int
		lastvisit *TreeNode
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		node := stack[len(stack)-1]
		if node.right == nil || node.right == lastvisit {
			stack = stack[:len(stack)-1] // 弹出
			result = append(result, node.val)
			lastvisit = node //保存访问记录
		} else {
			root = node.right // 访问右节点
		}
	}
	fmt.Println(result)
}
