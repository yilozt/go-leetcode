/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
package main

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		count += int(num & 1)
		num >>= 1
	}
	return count
}

// @lc code=end
