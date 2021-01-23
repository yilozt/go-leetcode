/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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

func reversed(arg [][]int) [][]int {
	for i, j := 0, len(arg)-1; i < j; i, j = i+1, j-1 {
		arg[i], arg[j] = arg[j], arg[i]
	}
	return arg
}

func levelOrderBottom(root *TreeNode) [][]int {
	var (
		result [][]int
		queue  []*TreeNode
		length int
	)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		var tmp []int
		length = len(queue)
		for i := 0; i < length; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, tmp)
		queue = queue[length:]
	}
	return reversed(result)
}

// @lc code=end
