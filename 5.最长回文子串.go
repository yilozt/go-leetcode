/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
package main

func longestPalindrome(s string) string {
	var result string
	for i := range s {
		sideOdd := extend(s, i, i)
		sideEven := extend(s, i, i+1)
		if sideOdd[1]-sideOdd[0] > len(result) {
			result = s[sideOdd[0]:sideOdd[1]]
		}
		if sideEven[1]-sideEven[0] > len(result) {
			result = s[sideEven[0]:sideEven[1]]
		}
	}
	return result
}

func extend(s string, l, r int) [2]int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return [2]int{l + 1, r}
}

// @lc code=end
