/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

package main

// @lc code=start
func maxArea(height []int) int {
	area := 0
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			area = max(area, (right-left)*height[left])
			left++
		} else {
			area = max(area, (right-left)*height[right])
			right--
		}
	}
	return area
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// @lc code=end
