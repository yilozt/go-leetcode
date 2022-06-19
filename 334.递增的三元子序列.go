/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

package main

import "math"

// @lc code=start
func increasingTriplet(nums []int) bool {
	first := nums[0]
	second := math.MaxInt

	for _, num := range nums {
		if num > second {
			return true
		}
		if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

// @lc code=end

func main() {
	increasingTriplet([]int{20, 100, 10, 12, 5, 13})
}
