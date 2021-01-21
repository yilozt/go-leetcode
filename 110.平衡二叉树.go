/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// 如果左子树失衡/右子树失衡/以自己为根的树失衡
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	// 否则正常处理返回高度
	if left > right {
		return left + 1
	}
	return right + 1
}

func abs(arg int) int {
	if arg >= 0 {
		return arg
	}
	return -arg
}

// @lc code=end
