/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
package main

func permute(nums []int) [][]int {
	var (
		result    [][]int
		tmp       []int
		visited   []bool
		backtrack func(int)
	)
	visited = make([]bool, len(nums))
	backtrack = func(pos int) {
		if len(tmp) == len(nums) {
			result = append(result, append([]int(nil), tmp...))
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				tmp = append(tmp, nums[i])
				visited[i] = true
				backtrack(i + 1)
				visited[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrack(0)
	return result
}

// @lc code=end
