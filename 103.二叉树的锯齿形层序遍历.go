/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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

func reverse(arg []int) {
	for i, j := 0, len(arg)-1; i < j; i, j = i+1, j-1 {
		arg[i], arg[j] = arg[j], arg[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var (
		result [][]int
		queue  []*TreeNode
		length int
		count  = false
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
		if count {
			reverse(tmp)
		}
		count = !count
		result = append(result, tmp)
		queue = queue[length:]
	}
	return result
}

// @lc code=end
