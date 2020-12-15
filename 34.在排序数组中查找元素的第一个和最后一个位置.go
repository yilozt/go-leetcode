/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package main

func searchRange(nums []int, target int) []int {
	num1, num2 := binarySearch(nums, target, true), binarySearch(nums, target, false)-1
	if num1 <= num2 && num2 < len(nums) && nums[num1] == target && nums[num2] == target {
		return []int{num1, num2}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, first bool) int {
	ans := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (first && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

// @lc code=end
