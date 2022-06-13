/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

package main

// @lc code=start
func containsDuplicate(nums []int) bool {
	count := make(map[int]int, 0)

	for _, num := range nums {
		count[num] += 1
		if count[num] >= 2 {
			return true
		}
	}
	return false
}

// @lc code=end
