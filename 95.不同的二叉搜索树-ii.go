/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var result []*TreeNode
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				tmp := &TreeNode{Val: i}
				tmp.Left = left[j]
				tmp.Right = right[k]
				result = append(result, tmp)
			}
		}
	}
	return result
}

// @lc code=end
