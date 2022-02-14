package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
	var result []string
	Serialize(test, &result)
	fmt.Println(result)
	var rebuildRoot *TreeNode
	var index int
	DeSerialize(&rebuildRoot, result, &index)
	print(rebuildRoot)
}

func print(root *TreeNode) {
	var list []*TreeNode
	list = append(list, root)
	for len(list) != 0 {
		length := len(list)
		for _, node := range list[:length] {
			if node != nil {
				fmt.Print(node.Val, " ")
				list = append(list, node.Left)
				list = append(list, node.Right)
			} else {
				fmt.Print("$ ")
			}
		}
		list = list[length:]
		fmt.Println()
	}
}

func Serialize(root *TreeNode, result *[]string) {
	if root == nil {
		*result = append(*result, "$")
		return
	}
	*result = append(*result, strconv.Itoa(root.Val))
	Serialize(root.Left, result)
	Serialize(root.Right, result)
}

func DeSerialize(root **TreeNode, input []string, index *int) {
	if input[*index] != "$" {
		*root = &TreeNode{}
		num, _ := strconv.Atoi(input[*index])
		(*root).Val = num
		(*index)++
		DeSerialize(&(*root).Left, input, index)
		DeSerialize(&(*root).Right, input, index)
	} else {
		(*index)++
	}
}
