/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

package main

// @lc code=start
func maxSubArray(nums []int) int {
	prev := 0
	maxAns := nums[0]

	for _, num := range nums {
		prev = max(num, prev+num)
		maxAns = max(prev, maxAns)
	}

	return maxAns
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
