/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

func levelOrder(root *TreeNode) [][]int {
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
	return result
}

// @lc code=end
