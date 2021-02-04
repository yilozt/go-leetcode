/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
package main

import (
	"strings"
)

// 存储阶乘
var factorial []int

func calculateFactorial(n int) {
	factorial = make([]int, n)
	tmp := 1
	for i := 1; i <= n; i++ {
		tmp *= i
		factorial[i-1] = tmp
	}
}

func getPermutation(n int, k int) string {
	// 构建数字集合
	nums := make([]byte, n)
	for i := range nums {
		nums[i] = byte(i + 1 + '0')
	}
	// 计算阶乘
	calculateFactorial(n - 1)
	// 储存结果
	var result strings.Builder
	k -= 1
	for n >= 2 {
		index := k / factorial[n-2]
		result.WriteByte(nums[index])
		k = k % factorial[n-2]
		n--
		nums = append(nums[:index], nums[index+1:]...)
	}
	result.Write(nums)
	return result.String()
}

// @lc code=end
