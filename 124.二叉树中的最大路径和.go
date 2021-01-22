/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

import "math"

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, result)
	right := helper(root.Right, result)
	*result = max(*result, left+right+root.Val)
	return max(max(left+root.Val, right+root.Val), 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
