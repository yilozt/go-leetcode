/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

package main

// @lc code=start
func hammingWeight(num uint32) int {
	count := uint32(0)
	for num != 0 {
		count += (num & 1)

		num >>= 1
	}
	return int(count)
}

// @lc code=end
