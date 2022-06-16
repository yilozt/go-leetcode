/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

package main

// f(n) = f(n-1) + f(n-2)
// 使用循环
func climbStairs(n int) int {
	prev := 1
	cur := 2

	if n < 2 {
		return prev
	}

	for i := 0; i < n-2; i++ {
		tmp := cur
		cur += prev
		prev = tmp
	}

	return cur
}

// @lc code=end
