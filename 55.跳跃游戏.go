/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
package main

func canJump(nums []int) bool {
	var farthest int
	for i := 0; i < len(nums); i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
		}
	}
	if farthest >= len(nums)-1 {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
