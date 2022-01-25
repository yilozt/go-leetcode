package main

// 动态规划

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		if i != n { // 在该段绳子为中间过程时可以预设最大乘积为本身，否则因为绳子至少切成两段，不能做此预设
			dp[i] = i
		}
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
