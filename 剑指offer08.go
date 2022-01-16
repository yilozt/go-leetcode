package main

type TreeNode struct {
	left, right *TreeNode
	father      *TreeNode
}

func GetNext(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		tmp := node.right
		for tmp.left != nil {
			tmp = tmp.left
		}
		return tmp
	}
	tmp := node
	for tmp.father != nil && tmp.father.right == tmp {
		tmp = tmp.father
	}
	return tmp.father
}
