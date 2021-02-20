/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	cache := make(map[string]bool)
	for _, v := range wordDict {
		cache[v] = true
	}
	for i := 0; i < len(s); i++ {
		if _, ok := cache[s[:(i+1)]]; ok {
			dp[i] = true
		} else {
			for j := i; j > 0; j-- {
				if _, ok := cache[s[j:(i+1)]]; ok && dp[j-1] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end
