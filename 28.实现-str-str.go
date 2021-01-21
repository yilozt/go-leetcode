/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
package main

func strStr(haystack string, needle string) int {
	var i, j int
	for i = 0; i <= len(haystack)-len(needle); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}

// @lc code=end
