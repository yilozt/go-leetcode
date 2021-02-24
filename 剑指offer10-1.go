package main

func fib(n int) int {
	// 简单动态规划
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		if dp[i] > (1e9 + 7) {
			dp[i] -= (1e9 + 7)
		}
	}
	return dp[n]
}
