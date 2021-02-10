/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
package main

func reverseBits(num uint32) uint32 {
	var bit, result uint32
	for i := 0; i < 32; i++ {
		result <<= 1
		bit = num & 1
		result = result + bit
		num >>= 1
	}
	return result
}

// @lc code=end
