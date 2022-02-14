package main

func isMatch(s string, p string) bool {
	matches := func(i, j int) bool {
		// matches函数用于判断s的第i位和p的第j位是否相同
		// s的第零位为空，所以直接返回false
		if i == 0 {
			return false
		}
		// .是一定可以匹配上的
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	dp := make([][]bool, len(s)+1) // dp[i][j]的实际含义是s的第i位和p的第j位是否匹配
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				// 遇到*说明可以匹配多个前一位字符，也可以不匹配
				dp[i][j] = dp[i][j] || dp[i][j-2]
				// 匹配成功，结果对dp[i-1][j]取或
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				// 纯字母匹配成功
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
