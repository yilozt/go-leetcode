/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
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
