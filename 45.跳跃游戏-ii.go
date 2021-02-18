/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
package main

func jump(nums []int) int {
	var maxPosition, end, step int
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step += 1
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
