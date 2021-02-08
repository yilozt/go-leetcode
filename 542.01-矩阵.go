/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
package main

func updateMatrix(matrix [][]int) [][]int {
	// 本质是一个广度优先搜索
	var queue [][2]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}
	direction := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) != 0 {
		coor := queue[0]
		queue = queue[1:]
		for _, dir := range direction {
			i := coor[0] + dir[0]
			j := coor[1] + dir[1]
			if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && matrix[i][j] == -1 {
				matrix[i][j] = matrix[coor[0]][coor[1]] + 1
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	return matrix
}

// @lc code=end
