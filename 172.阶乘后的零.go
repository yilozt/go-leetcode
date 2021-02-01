/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
package main

func trailingZeroes(n int) int {
	base := 5
	count := 0
	for n >= base {
		count += n / base
		base *= 5
	}
	return count
}

// @lc code=end
