/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
package main

import "sort"

func permuteUnique(nums []int) [][]int {
	var (
		result    [][]int
		tmp       []int
		visited   []bool
		backtrack func(int)
	)
	visited = make([]bool, len(nums))
	sort.Ints(nums)
	backtrack = func(pos int) {
		if len(tmp) == len(nums) {
			result = append(result, append([]int(nil), tmp...))
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			tmp = append(tmp, nums[i])
			visited[i] = true
			backtrack(i + 1)
			visited[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(0)
	return result
}

// @lc code=end
