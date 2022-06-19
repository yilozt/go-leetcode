/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sum(root *TreeNode, count map[int]int) int {
	if root == nil {
		return 0
	}
	s := root.Val + sum(root.Left, count) + sum(root.Right, count)
	count[s]++
	return s
}

func findFrequentTreeSum(root *TreeNode) []int {
	count := make(map[int]int, 0)
	sum(root, count)

	ans := make([]int, 0)

	max := 0
	for k, c := range count {
		if c > max {
			max = c
			ans = make([]int, 0)
		}
		if c == max {
			ans = append(ans, k)
		}
	}

	return ans
}

// @lc code=end
