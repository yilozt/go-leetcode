/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
package main

func findAnagrams(s string, p string) []int {
	var result []int
	needs, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}
	left, right := 0, 0
	match := 0
	for right < len(s) {
		c := s[right]
		right++
		if needs[c] != 0 {
			window[c]++
			if needs[c] == window[c] {
				match += 1
			}
		}
		if right-left == len(p) {
			if match == len(needs) {
				result = append(result, left)
			}
			c := s[left]
			left++
			if needs[c] != 0 {
				if needs[c] == window[c] {
					match--
				}
				window[c]--
			}
		}
	}
	return result
}

// @lc code=end
