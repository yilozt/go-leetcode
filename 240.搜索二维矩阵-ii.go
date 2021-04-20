/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 || len(matrix[0]) == 0 {
		return false
	}
	columns := len(matrix[0])
	i, j := rows-1, 0
	for j < columns && i >= 0 {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
