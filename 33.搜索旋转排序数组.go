/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

package main

// @lc code=start
func search(nums []int, target int) int {
	off := getOffset(nums)
	return searchWithOffset(nums, target, off)
}

func getOffset(nums []int) int {
	l, r := 0, len(nums)-1

	for nums[l] > nums[r] {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
			continue
		}
		if nums[l] > nums[m] {
			r = m
			continue
		}
		return l
	}

	return l
}

func searchWithOffset(nums []int, target int, off int) int {
	l, r := 0, len(nums)-1

	addr := func(i int) int {
		return (i + off) % len(nums)
	}

	for l <= r {
		m := (l + r) / 2
		if nums[addr(m)] == target {
			return addr(m)
		}
		if target < nums[addr(m)] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

// @lc code=end
