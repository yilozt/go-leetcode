/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
package main

import (
	"math"
)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	tmp := make([][]int, m)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	count := 1
	// 构造矩阵
	for i := 0; i < m; i++ {
		count = 1
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				tmp[i][j] = count
				count++
			} else {
				count = 1
			}
		}
	}
	finalResult := math.MinInt32
	// 对每列使用单调栈进行求解
	for j := 0; j < n; j++ {
		var stack []int
		// 每行元素指定宽的上下边界
		top, bottom := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			for len(stack) != 0 && tmp[stack[len(stack)-1]][j] > tmp[i][j] {
				bottom[stack[len(stack)-1]] = i - 1
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				top[i] = -1
			} else {
				top[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		for _, item := range stack {
			bottom[item] = m - 1
		}
		for i := 0; i < m; i++ {
			finalResult = max(finalResult, (bottom[i]-top[i])*tmp[i][j])
		}
	}
	return finalResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
