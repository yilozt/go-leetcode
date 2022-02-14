package main

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func mirrorTree(root *TreeNode) *TreeNode { // 循环式写法事实上更加简单
	var list []*TreeNode
	list = append(list, root)
	for len(list) != 0 {
		length := len(list)
		for _, node := range list[:length] {
			if node != nil {
				node.Right, node.Left = node.Left, node.Right
				list = append(list, node.Right)
				list = append(list, node.Left)
			}
		}
		list = list[length:]
	}
	return root
}
