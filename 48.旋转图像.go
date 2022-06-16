/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

package main

import (
	"fmt"
	"reflect"
)

// @lc code=start
func rotate(matrix [][]int) {
	len := len(matrix)

	// 水平翻转
	for i := 0; i < len/2; i++ {
		for j := 0; j < len; j++ {
			swap(&matrix[i][j], &matrix[len-1-i][j])
		}
	}

	// 对角线翻转
	for i := 0; i < len; i++ {
		for j := 0; j < i; j++ {
			swap(&matrix[i][j], &matrix[j][i])
		}
	}
}

func swap(a *int, b *int) {
	tmp := (*a)
	(*a) = (*b)
	(*b) = tmp
}

// @lc code=end

func main() {
	orig := [][]int{
		{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16},
	}

	rotated := make([][]int, len(orig))
	copy(rotated, orig)

	rotate(rotated)

	fmt.Println("orig    : ", orig)
	fmt.Println("rotated : ", rotated)

	if reflect.DeepEqual(orig, rotated) {
		fmt.Println("OK!")
	}
}
