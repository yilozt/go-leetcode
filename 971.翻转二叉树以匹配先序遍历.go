/*
 * @lc app=leetcode.cn id=971 lang=golang
 *
 * [971] 翻转二叉树以匹配先序遍历
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

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var result []int
	index := 0
	helper(root, voyage, &result, &index)
	return result
}

func helper(root *TreeNode, voyage []int, result *[]int, index *int) {
	if root == nil {
		return
	}
	if root.Val != voyage[*index] {
		*result = []int{-1}
		return
	}
	if root.Left != nil && root.Left.Val != voyage[*index+1] {
		*result = append(*result, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}
	*index += 1
	helper(root.Left, voyage, result, index)
	helper(root.Right, voyage, result, index)
}

// @lc code=end
