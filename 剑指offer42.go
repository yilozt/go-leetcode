package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)
	result := math.MinInt32
	for i := 1; i <= len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i-1], nums[i-1])
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
