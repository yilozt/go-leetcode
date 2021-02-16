/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
package main

func minimumTotal(triangle [][]int) int {
	cache := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		cache[i] = make([]int, len(triangle[i]))
	}
	// 最后一层为基准值
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		cache[len(triangle)-1][j] = triangle[len(triangle)-1][j]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			cache[i][j] = min(cache[i+1][j], cache[i+1][j+1]) + triangle[i][j]
		}
	}
	return cache[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
