/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}

// @lc code=end
