/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
package main

func sumSubarrayMins(arr []int) int {
	// 首先使用单调栈求出某个数作为最小值的区间
	length := len(arr)
	var stack []int
	left, right := make([]int, length), make([]int, length)
	for index, value := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > value {
			right[stack[len(stack)-1]] = index
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[index] = -1
		} else {
			left[index] = stack[len(stack)-1]
		}
		stack = append(stack, index)
	}
	for _, value := range stack {
		right[value] = len(arr)
	}
	// 接着计算在每个区间内具有的子数组的个数，乘以这个区间的最小值
	var result int64
	for index, value := range arr {
		result += int64((index - left[index]) * (right[index] - index) * value)
	}
	return int(result % 1000000007)
}

// @lc code=end
