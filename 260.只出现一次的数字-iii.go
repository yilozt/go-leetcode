/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
package main

func singleNumber(nums []int) []int {
	var diff int
	for _, value := range nums {
		diff ^= value
	}
	// 设结果为a、b,当前得到的diff是a^b的结果
	// 取到diff的任意一位1（此处使用的是最后一位）
	diff = (diff & (diff - 1)) ^ diff
	result := []int{0, 0}
	for _, value := range nums {
		if value&diff == 0 {
			result[0] ^= value
		} else {
			result[1] ^= value
		}
	}
	return result
}

// @lc code=end
