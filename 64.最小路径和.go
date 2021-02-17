/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
package main

func minPathSum(grid [][]int) int {
	// 初始化dp数组
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	// 写入基准值
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	// 进行dp
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
