/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	target := nums[len(nums)-1]
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			left = mid
		} else {
			right = mid
		}
	}
	return min(nums[left], nums[right])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
