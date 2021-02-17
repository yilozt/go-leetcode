/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid)+1)
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0])+1)
	}
	// 这样可以确保起点dp[1][1]的路径数是1
	dp[0][1] = 1
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)][len(obstacleGrid[0])]
}

// @lc code=end
