/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var result []int
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for _, node := range queue[:length] {
			result = append(result, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
	}
	return result
}
