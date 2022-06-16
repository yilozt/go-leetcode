/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

package main

import "fmt"

// @lc code=start
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)

	// 只有一行的情况 (rowIndex = 0)
	res[0] = 1

	for row := 1; row <= rowIndex; row++ {
		// 先填充两侧数字
		res[0] = 1
		res[row] = 1

		// 根据上一次迭代的结果，计算中间数字
		prev := res[0]
		cur := res[1]
		for col := 1; col < row; col++ {
			res[col] = prev + cur
			prev = cur
			cur = res[col+1]
		}
	}

	return res
}

// @lc code=end

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(5))
}
