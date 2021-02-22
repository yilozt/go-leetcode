/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package main

func subsets(nums []int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}
		tmp = append(tmp, nums[cur])
		dfs(cur + 1)
		tmp = tmp[:len(tmp)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return res
}

// @lc code=end
