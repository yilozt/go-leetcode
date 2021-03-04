/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 保存临时变量
	left, right := root.Left, root.Right
	root.Left = mirrorTree(right)
	root.Right = mirrorTree(left)
	return root
}
