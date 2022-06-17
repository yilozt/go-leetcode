/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

package main

import "fmt"

// @lc code=start
func integerBreak(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	// n 大于 4 时开始划分
	products := make([]int, n+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
		}
		products[i] = max
	}

	return products[n]
}

// @lc code=end

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}
