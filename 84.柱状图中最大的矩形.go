/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
package main

import (
	"math"
)

func largestRectangleArea(heights []int) int {
	// 使用单调栈求每个高的左右边界
	left, right := make([]int, len(heights)), make([]int, len(heights))
	var stack []int
	for index, value := range heights {
		for len(stack) != 0 && heights[stack[len(stack)-1]] > value {
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
		right[value] = len(heights)
	}
	result := math.MinInt32
	for index, value := range heights {
		result = max(result, (right[index]-left[index]-1)*value)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
