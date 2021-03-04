/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var result [][]int
	var flag bool
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		var tmp []int
		for _, node := range queue[:length] {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
		if flag {
			reverseSlice(tmp)
		}
		flag = !flag
		result = append(result, tmp)
	}
	return result
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
