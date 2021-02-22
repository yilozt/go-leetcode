/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
package main

import "math"

func lengthOfLongestSubstring(s string) int {
	result := math.MinInt32
	left, right := 0, 0
	window := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		right++
		// 已经出现过
		if window[c] == 1 {
			result = max(right-left-1, result)
			for s[left] != c {
				char := s[left]
				left++
				window[char]--
			}
			left++
		} else {
			window[c]++
		}
	}
	result = max(right-left, result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
