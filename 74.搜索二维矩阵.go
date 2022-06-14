/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)*len(matrix[0])-1
	for l <= r {
		m := (l + r) / 2
		if target == index(matrix, m) {
			return true
		}
		if target < index(matrix, m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}

func index(matrix [][]int, i int) int {
	y := i / len(matrix[0])
	x := i % len(matrix[0])

	return matrix[y][x]
}

// @lc code=end
