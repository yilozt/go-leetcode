/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var (
		result [][]int
		tmp    []int
	)
	sort.Ints(nums)
	backtrack(&result, tmp, nums, 0)
	return result
}

func backtrack(result *[][]int, tmp, nums []int, pos int) {
	*result = append(*result, append([]int{}, tmp...))
	for i := pos; i < len(nums); i++ {
		if i > pos && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		backtrack(result, tmp, nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

// @lc code=end
