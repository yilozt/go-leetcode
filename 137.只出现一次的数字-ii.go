/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
package main

func singleNumber(nums []int) int {
	var a, b int
	for _, value := range nums {
		b = (b ^ value) & ^a
		a = (a ^ value) & ^b
	}
	return b
}

// @lc code=end
