/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start
package main

func rangeBitwiseAnd(m int, n int) int {
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}

// @lc code=end
