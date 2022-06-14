/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

package main

// @lc code=start
func majorityElement(nums []int) int {
	m := make(map[int]int, 0)
	len := len(nums)
	for _, n := range nums {
		m[n] += 1
		if m[n] > len/2 {
			return n
		}
	}
	return -1
}

// @lc code=end
