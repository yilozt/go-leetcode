/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(left, right *TreeNode) bool {
	// 均为nil
	if left == nil && right == nil {
		return true
	}
	// 有一边为nil
	if left == nil || right == nil {
		return false
	}
	// 全不为nil
	return left.Val == right.Val && isSame(left.Right, right.Left) && isSame(left.Left, right.Right)
}
