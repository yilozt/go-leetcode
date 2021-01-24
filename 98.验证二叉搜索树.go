/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

type result struct {
	isValid        bool
	maxVal, minVal *TreeNode
}

func isValidBST(root *TreeNode) bool {
	result := helper(root)
	return result.isValid
}

func helper(root *TreeNode) result {
	res := result{}
	if root == nil {
		res.isValid = true
		return res
	}
	left := helper(root.Left)
	right := helper(root.Right)
	// 如果有至少一边不是二叉搜索树
	if !left.isValid || !right.isValid {
		return res
	}
	// 左子树最大值大于等于根节点
	if left.maxVal != nil && left.maxVal.Val >= root.Val {
		return res
	}
	// 右子树最小值小于等于根节点
	if right.minVal != nil && right.minVal.Val <= root.Val {
		return res
	}
	res.isValid = true
	res.minVal = root
	if left.minVal != nil {
		res.minVal = left.minVal
	}
	res.maxVal = root
	if right.maxVal != nil {
		res.maxVal = right.maxVal
	}
	return res
}

// @lc code=end
