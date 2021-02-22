/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
package main

func checkInclusion(s1 string, s2 string) bool {
	window, needs := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		needs[s1[i]]++
	}
	left, right := 0, 0
	match := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if needs[c] != 0 {
			window[c]++
			if needs[c] == window[c] {
				match += 1
			}
		}
		for match == len(needs) {
			if right-left == len(s1) {
				return true
			}
			c := s2[left]
			left++
			if window[c] != 0 {
				if window[c] == needs[c] {
					match -= 1
				}
				window[c] -= 1
			}
		}
	}
	return false
}

// @lc code=end
