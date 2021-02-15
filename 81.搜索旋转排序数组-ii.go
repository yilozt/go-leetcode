/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
package main

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left+1 < right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] < nums[mid] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid
			} else {
				left = mid
			}
		} else if nums[right] > nums[mid] {
			if nums[right] >= target && nums[mid] <= target {
				left = mid
			} else {
				right = mid
			}
		}
	}
	if nums[left] == target || nums[right] == target {
		return true
	}
	return false
}

// @lc code=end
