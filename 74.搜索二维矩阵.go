/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
package main

func searchMatrix(matrix [][]int, target int) bool {
	// 本质是一个带地址映射的二分查找
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	left, right := 0, len(matrix)*len(matrix[0])-1
	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/len(matrix[0])][mid%len(matrix[0])]
		if midValue > target {
			right = mid - 1
		} else if midValue < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
