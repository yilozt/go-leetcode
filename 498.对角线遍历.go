/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

package main

import (
	"fmt"
	"reflect"
)

// @lc code=start

type Pos struct {
	x int
	y int
}

func findDiagonalOrder(mat [][]int) []int {
	// 移动方向
	dir_up_right := true
	// 当前位置
	pos := Pos{x: 0, y: 0}
	// 矩阵大小
	x_size := len(mat[0])
	y_size := len(mat)

	res := make([]int, 0)

	for pos.x < x_size && pos.y < y_size {
		res = append(res, mat[pos.y][pos.x])

		next_pos := Pos{x: 0, y: 0}
		should_change_dir := false

		if dir_up_right {
			next_pos = Pos{x: pos.x + 1, y: pos.y - 1}
		} else {
			next_pos = Pos{x: pos.x - 1, y: pos.y + 1}
		}

		// 检查边界，并修正位置
		if next_pos.y < 0 {
			next_pos.x = pos.x + 1
			next_pos.y = 0
			should_change_dir = true
		}
		if next_pos.x < 0 {
			next_pos.x = 0
			next_pos.y = pos.y + 1
			should_change_dir = true
		}
		if next_pos.x == x_size {
			next_pos.x = x_size - 1
			next_pos.y = pos.y + 1
			should_change_dir = true
		}
		if next_pos.y == y_size {
			next_pos.x = pos.x + 1
			next_pos.y = y_size - 1
			should_change_dir = true
		}

		if should_change_dir {
			dir_up_right = !dir_up_right
		}

		pos = next_pos
	}

	return res
}

// @lc code=end

func main() {
	datas := []struct {
		data [][]int
		ans  []int
	}{
		{
			data: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			ans:  []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			data: [][]int{{1, 2}, {3, 4}},
			ans:  []int{1, 2, 3, 4},
		},
	}

	for _, data := range datas {
		got := findDiagonalOrder(data.data)
		want := data.ans
		if !reflect.DeepEqual(got, want) {
			fmt.Printf("got: %v, want: %v\n", got, want)
			return
		}
	}
}
