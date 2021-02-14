/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
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
	if nums[left] == target {
		return left
	} else if nums[right] == target {
		return right
	}
	return -1
}

// @lc code=end
