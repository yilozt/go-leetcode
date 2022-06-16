/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

package main

// @lc code=start
func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2

		// 无法判断该向哪个方向划分，直接顺序查找
		if nums[l] == nums[m] && nums[m] == nums[r] {
			for i := l; i <= r; i++ {
				if nums[i] < nums[l] {
					return nums[i]
				}
			}
			return nums[l]
		}

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[r]
}

// @lc code=end

func main() {
	findMin([]int{3, 3, 1, 3})
}
