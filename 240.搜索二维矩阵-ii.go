/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

package main

import "fmt"

// @lc code=start
func search(matrix [][]int, target, col_start, row_start, col_end, row_end int) bool {
	if col_start > col_end || row_start > row_end {
		return false
	}

	col_mid := (col_end + col_start) / 2
	row_mid := (row_end + row_start) / 2
	mid := matrix[row_mid][col_mid]

	if mid == target {
		return true
	} else if target < mid {
		return search(matrix, target, col_start, row_start, col_mid-1, row_end) ||
			search(matrix, target, col_mid, row_start, col_end, row_mid-1)
	} else {
		return search(matrix, target, col_mid+1, row_start, col_end, row_end) ||
			search(matrix, target, col_start, row_mid+1, col_mid, row_end)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return search(matrix, target, 0, 0, len(matrix[0])-1, len(matrix)-1)
}

// @lc code=end

func test(mat [][]int, target int, res bool) {
	if res != searchMatrix(mat, target) {
		fmt.Printf("FAILD: target: %d, mat: %v\n", target, mat)
	}
}

func main() {
	mat := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5
	test(mat, target, true)

	target = -1
	test(mat, target, false)

	target = 20
	test(mat, target, false)

	target = 100
	test(mat, target, false)

	mat = [][]int{{1, 4, 7, 11, 15}}
	target = 4
	test(mat, target, true)

	target = 16
	test(mat, target, false)

	mat = [][]int{{1}, {4}, {7}, {11}}
	target = 4
	test(mat, target, true)

	target = 16
	test(mat, target, false)

	mat = [][]int{{4}}
	target = 4
	test(mat, target, true)

	target = -1
	test(mat, target, false)

	mat = [][]int{{5}, {10}}
	target = 5
	test(mat, target, true)

	mat = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	target = 5
	test(mat, target, true)
}
