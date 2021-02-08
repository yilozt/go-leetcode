/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
package main

func singleNumber(nums []int) int {
	result := 0
	for _, value := range nums {
		result ^= value
	}
	return result
}

// @lc code=end
