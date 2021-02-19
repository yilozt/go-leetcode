/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
package main

var (
	str  string
	flag [][]bool
)

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s))
	str = s
	flag = make([][]bool, len(s))
	for i := range flag {
		flag[i] = make([]bool, len(s))
	}
	isPalindromic(len(s))
	for i := 0; i < len(s); i++ {
		if flag[0][i] {
			dp[i] = 0
			continue
		}
		dp[i] = i
		for j := 1; j <= i; j++ {
			if flag[j][i] {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	return dp[len(s)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func isPalindromic(length int) {
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if (str[i] == str[j]) && (j-i <= 1 || flag[i+1][j-1]) {
				flag[i][j] = true
			}
		}
	}
}

// @lc code=end
