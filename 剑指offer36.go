package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Convert(root *TreeNode) *TreeNode { // 给定二叉搜索树的根节点，将其转换为双向排序链表
	var lastNodeInList *TreeNode = nil
	ConvertNode(root, &lastNodeInList)
	if lastNodeInList == nil {
		return nil
	}
	for lastNodeInList.Left != nil {
		lastNodeInList = lastNodeInList.Left
	}
	return lastNodeInList
}

func ConvertNode(root *TreeNode, lastNodeInList **TreeNode) {
	if root == nil {
		return
	}
	ConvertNode(root.Left, lastNodeInList)
	root.Left = *lastNodeInList
	if *lastNodeInList != nil {
		(*lastNodeInList).Right = root
	}
	*lastNodeInList = root
	ConvertNode(root.Right, lastNodeInList)
}

func main() {
	test := &TreeNode{
		10,
		&TreeNode{
			6,
			&TreeNode{4, nil, nil},
			&TreeNode{8, nil, nil},
		},
		&TreeNode{
			14,
			&TreeNode{12, nil, nil},
			&TreeNode{16, nil, nil},
		},
	}
	node := Convert(test)
	if node == nil {
		return
	}
	for node.Right != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Right
	}
	fmt.Printf("%d.", node.Val)
	fmt.Println()
	for node.Left != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Left
	}
	fmt.Printf("%d.", node.Val)
}
