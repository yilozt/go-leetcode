/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
package main

import "math"

func minWindow(s string, t string) string {
	window := make(map[byte]int)
	needs := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	left, right := 0, 0
	start, end := 0, 0
	min := math.MaxInt32
	match := 0
	for right < len(s) {
		c := s[right]
		right += 1
		if needs[c] != 0 {
			window[c]++
			if window[c] == needs[c] {
				match++
			}
		}
		for match == len(needs) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left += 1
			if needs[c] != 0 {
				if window[c] == needs[c] {
					match -= 1
				}
				window[c] -= 1
			}
		}
	}
	if min == math.MaxInt32 {
		return ""
	}
	return s[start:end]
}

// @lc code=end
